package app

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/takke/dorm-laundry/internal/laundry"
)

type Server struct {
	router chi.Router
	store  *laundry.Store
	root   string
}

type pageData struct {
	Title                  string
	CurrentPath            string
	Machines               []laundry.Machine
	Machine                laundry.Machine
	MachineSummary         []laundry.StatusSummary
	CurrentSession         *laundry.Session
	NotificationPreference laundry.NotificationPreference
	FlashMessage           string
	ErrorMessage           string
	Now                    time.Time
}

var (
	iconAssetPaths = map[string]string{
		"washer":   "smart-wash-smart-clean-smart-cleaner-washing-machine-laundry-smart-loundry-svgrepo-com.svg",
		"timer":    "timer-svgrepo-com.svg",
		"basket":   "basket-svgrepo-com.svg",
		"bell":     "bell-alt-svgrepo-com.svg",
		"settings": "setting-5-svgrepo-com.svg",
		"check":    "check-read-svgrepo-com.svg",
	}
	iconAssetCache sync.Map
)

func NewServer() (*Server, error) {
	server := &Server{
		store: laundry.NewStore(),
		root:  projectRoot(),
	}
	server.router = server.routes()

	return server, nil
}

func (s *Server) Handler() http.Handler {
	return s.router
}

func parseTemplates(patterns ...string) (*template.Template, error) {
	funcs := template.FuncMap{
		"statusLabel":       laundry.StatusLabel,
		"statusDescription": laundry.StatusDescription,
		"statusSummary":     laundry.MachineStatusSummary,
		"actionHint":        laundry.ActionHint,
		"relativeMinutes":   laundry.RelativeMinutes,
		"countByStatus":     laundry.CountByStatus,
		"sessionProgressAt": func(session *laundry.Session, now time.Time) int {
			if session == nil {
				return 0
			}
			return laundry.SessionProgressPercent(*session, now)
		},
		"sessionStatusLabel":    laundry.SessionStatusLabel,
		"notificationStateText": notificationStateText,
		"sessionSummaryAt": func(session *laundry.Session, now time.Time) string {
			if session == nil {
				return ""
			}
			return laundry.SessionSummary(*session, now)
		},
		"formatDateTime": func(value string) string {
			date, err := time.Parse(time.RFC3339, value)
			if err != nil {
				return value
			}
			return date.Format("1/2 15:04")
		},
		"formatClock": func(value string) string {
			date, err := time.Parse(time.RFC3339, value)
			if err != nil {
				return value
			}
			return date.Format("15:04")
		},
		"sessionRemainingMinutesAt": func(session *laundry.Session, now time.Time) int {
			if session == nil {
				return 0
			}
			endTime, err := time.Parse(time.RFC3339, session.EndsAt)
			if err != nil {
				return 0
			}
			remaining := int(endTime.Sub(now).Minutes())
			if remaining < 0 {
				return 0
			}
			return remaining
		},
		"progressCircleOffsetAt": func(session *laundry.Session, now time.Time) int {
			progress := 0
			if session != nil {
				progress = laundry.SessionProgressPercent(*session, now)
			}
			const circumference = 691
			return circumference - (progress * circumference / 100)
		},
		"statusBadgeClass": func(status any) string {
			switch laundry.MachineStatus(fmt.Sprint(status)) {
			case laundry.StatusAvailable:
				return "bg-status-available-soft text-status-available"
			case laundry.StatusRunning:
				return "bg-status-running-soft text-status-running"
			case laundry.StatusAwaitingPickup:
				return "bg-status-pickup-soft text-status-pickup"
			default:
				return "bg-surface-container-high text-outline"
			}
		},
		"machineCardToneClass": machineCardToneClass,
		"machineMetricClass":   machineMetricClass,
		"iconSVG":              iconSVG,
		"isCurrentPath": func(currentPath, href string) bool {
			return currentPath == href
		},
		"firstAvailableMachineID": func(machines []laundry.Machine) string {
			for _, machine := range machines {
				if machine.Status == laundry.StatusAvailable {
					return machine.ID
				}
			}
			return ""
		},
	}

	tmpl := template.New("base").Funcs(funcs)

	for _, pattern := range patterns {
		var err error
		tmpl, err = tmpl.ParseGlob(pattern)
		if err != nil {
			return nil, err
		}
	}

	return tmpl, nil
}

func projectRoot() string {
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		return filepath.Clean(filepath.Join(filepath.Dir(filename), "..", ".."))
	}

	dir, err := os.Getwd()
	if err != nil {
		return "."
	}

	return dir
}

func notificationStateText(preference laundry.NotificationPreference) string {
	if preference.Enabled {
		return "有効"
	}
	if !preference.Supported {
		return "未対応"
	}
	return "未設定"
}

func machineCardToneClass(status laundry.MachineStatus) string {
	switch status {
	case laundry.StatusAvailable:
		return "bg-status-available-soft text-status-available"
	case laundry.StatusRunning:
		return "bg-status-running-soft text-status-running"
	case laundry.StatusAwaitingPickup:
		return "bg-status-pickup-soft text-status-pickup"
	default:
		return "bg-surface-container-highest text-outline"
	}
}

func machineMetricClass(status laundry.MachineStatus) string {
	switch status {
	case laundry.StatusAvailable:
		return "text-status-available"
	case laundry.StatusRunning:
		return "text-accent"
	case laundry.StatusAwaitingPickup:
		return "text-status-pickup"
	default:
		return "text-outline"
	}
}

func iconSVG(name, class string) template.HTML {
	class = template.HTMLEscapeString(class)

	if svg, ok := fileBackedIconSVG(name, class); ok {
		return template.HTML(svg)
	}

	var body string
	switch name {
	case "location":
		body = `<path d="M12 21s-6-5.33-6-11a6 6 0 1 1 12 0c0 5.67-6 11-6 11Z"/><circle cx="12" cy="10" r="2.25" fill="currentColor"/>`
	case "washer":
		body = `<rect x="5" y="4" width="14" height="16" rx="4"/><circle cx="12" cy="13" r="3.75" fill="none" stroke="currentColor" stroke-width="1.8"/><path d="M9 7.75h.01M12 7.75h.01M15 7.75h.01" stroke="currentColor" stroke-width="2" stroke-linecap="round"/>`
	case "timer":
		body = `<circle cx="12" cy="13" r="6.5"/><path d="M12 13V9.5m0 3.5 2.5 2.5M9 3h6M15.5 5.5 17 4" stroke="currentColor" stroke-width="1.8" stroke-linecap="round" stroke-linejoin="round" fill="none"/>`
	case "basket":
		body = `<path d="M6.5 11.5h11l-1.2 6.3a2 2 0 0 1-2 1.6H9.7a2 2 0 0 1-2-1.6L6.5 11.5Z"/><path d="M9.5 11.5 12 7.5l2.5 4" fill="none" stroke="currentColor" stroke-width="1.8" stroke-linecap="round" stroke-linejoin="round"/>`
	case "wrench":
		body = `<path d="m14.5 6.5 3-3a3 3 0 0 1-3.9 3.9l-5.9 5.9a2 2 0 1 1-2.8-2.8l5.9-5.9A3 3 0 0 1 14.5 6.5Z"/>`
	case "home":
		body = `<path d="M4.5 10.5 12 4l7.5 6.5V19a1 1 0 0 1-1 1h-4.5v-5h-4v5H5.5a1 1 0 0 1-1-1z"/>`
	case "bell":
		body = `<path d="M12 4.5a4 4 0 0 1 4 4V11c0 .8.3 1.5.8 2.1l1 1.1c.3.3.1.8-.3.8H6.5c-.4 0-.6-.5-.3-.8l1-1.1c.5-.6.8-1.3.8-2.1V8.5a4 4 0 0 1 4-4Z"/><path d="M10 17a2 2 0 0 0 4 0" fill="none" stroke="currentColor" stroke-width="1.8" stroke-linecap="round"/>`
	case "settings":
		body = `<path d="M12 8.5A3.5 3.5 0 1 1 8.5 12 3.5 3.5 0 0 1 12 8.5Zm0-5 1.1 2.3 2.5.4.5 2.5 2.3 1.1-1.1 2.3 1.1 2.3-2.3 1.1-.5 2.5-2.5.4L12 20.5l-1.1-2.3-2.5-.4-.5-2.5-2.3-1.1L6.7 12 5.6 9.7l2.3-1.1.5-2.5 2.5-.4L12 3.5Z" fill="none" stroke="currentColor" stroke-width="1.5" stroke-linejoin="round"/>`
	case "plus":
		body = `<path d="M12 5v14M5 12h14" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"/>`
	case "check":
		body = `<circle cx="12" cy="12" r="8"/><path d="m8.5 12.5 2.3 2.3 4.7-5.1" fill="none" stroke="#fff" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/>`
	case "info":
		body = `<circle cx="12" cy="12" r="8"/><path d="M12 10.5v4m0-7h.01" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round"/>`
	case "drop":
		body = `<path d="M12 4.5c2.8 3.6 4.2 6.1 4.2 7.7A4.2 4.2 0 1 1 7.8 12c0-1.6 1.4-4.1 4.2-7.5Z"/>`
	case "thermo":
		body = `<path d="M10 6.5a2 2 0 1 1 4 0V13a3 3 0 1 1-4 0Z" fill="none" stroke="currentColor" stroke-width="1.8"/><path d="M12 10v5" fill="none" stroke="currentColor" stroke-width="1.8" stroke-linecap="round"/>`
	default:
		body = `<circle cx="12" cy="12" r="8"/>`
	}

	return template.HTML(fmt.Sprintf(
		`<svg viewBox="0 0 24 24" aria-hidden="true" class="%s" fill="currentColor" xmlns="http://www.w3.org/2000/svg">%s</svg>`,
		class,
		body,
	))
}

func fileBackedIconSVG(name, class string) (string, bool) {
	filename, ok := iconAssetPaths[name]
	if !ok {
		return "", false
	}

	raw, ok := iconAssetCache.Load(filename)
	if !ok {
		bytes, err := os.ReadFile(filepath.Join(projectRoot(), "web", "static", "icons", filename))
		if err != nil {
			return "", false
		}

		raw = string(bytes)
		iconAssetCache.Store(filename, raw)
	}

	svg := normalizeSVGMarkup(raw.(string), class)
	if svg == "" {
		return "", false
	}

	return svg, true
}

func normalizeSVGMarkup(raw, class string) string {
	start := strings.Index(raw, "<svg")
	if start == -1 {
		return ""
	}

	openEnd := strings.Index(raw[start:], ">")
	if openEnd == -1 {
		return ""
	}
	openEnd += start

	closeStart := strings.LastIndex(raw, "</svg>")
	if closeStart == -1 || closeStart <= openEnd {
		return ""
	}

	openTag := raw[start : openEnd+1]
	openTag = strings.Replace(openTag, "<svg", fmt.Sprintf(`<svg aria-hidden="true" class="%s"`, class), 1)

	replacer := strings.NewReplacer(
		` width="800px"`, "",
		` height="800px"`, "",
		` fill="#000000"`, ` fill="currentColor"`,
		` fill="#1C274C"`, ` fill="currentColor"`,
		` stroke="#000000"`, ` stroke="currentColor"`,
		` stroke="#323232"`, ` stroke="currentColor"`,
		` stroke="#1C274C"`, ` stroke="currentColor"`,
	)

	openTag = replacer.Replace(openTag)
	inner := replacer.Replace(raw[openEnd+1 : closeStart])

	return openTag + inner + "</svg>"
}

func (s *Server) routes() chi.Router {
	r := chi.NewRouter()

	fileServer(r, "/static", http.Dir(filepath.Join(s.root, "web", "static")))

	r.Get("/", s.handleHome)
	r.Get("/machines/{machineID}", s.handleMachineDetail)
	r.Get("/sessions/current", s.handleCurrentSession)

	r.Post("/sessions", s.handleStartSession)
	r.Post("/sessions/{sessionID}/claim", s.handleClaimSession)
	r.Post("/push-subscriptions", s.handleEnableNotifications)

	return r
}

func fileServer(r chi.Router, path string, root http.FileSystem) {
	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"
	r.Get(path, func(w http.ResponseWriter, req *http.Request) {
		http.StripPrefix(path[:len(path)-1], http.FileServer(root)).ServeHTTP(w, req)
	})
}

func (s *Server) handleHome(w http.ResponseWriter, r *http.Request) {
	machines := s.store.ListMachines()
	s.renderPage(w, "home", pageData{
		Title:          "洗濯機一覧",
		CurrentPath:    "/",
		Machines:       machines,
		MachineSummary: laundry.SummarizeMachines(machines),
		Now:            s.store.Now(),
	})
}

func (s *Server) handleMachineDetail(w http.ResponseWriter, r *http.Request) {
	machine, err := s.store.GetMachine(chi.URLParam(r, "machineID"))
	if err != nil {
		http.NotFound(w, r)
		return
	}

	s.renderPage(w, "machine_detail", pageData{
		Title:       machine.Name,
		CurrentPath: "/",
		Machine:     machine,
		Now:         s.store.Now(),
	})
}

func (s *Server) handleCurrentSession(w http.ResponseWriter, r *http.Request) {
	s.renderPage(w, "current_session", pageData{
		Title:                  "現在のセッション",
		CurrentPath:            "/sessions/current",
		CurrentSession:         s.store.CurrentSession(),
		NotificationPreference: s.store.NotificationPreference(),
		Now:                    s.store.Now(),
	})
}

func (s *Server) handleStartSession(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "invalid form", http.StatusBadRequest)
		return
	}

	machineID := r.FormValue("machine_id")
	machine, err := s.store.StartSession(machineID)

	data := pageData{
		Machine: machine,
		Now:     s.store.Now(),
	}

	if err != nil {
		if errors.Is(err, laundry.ErrMachineBusy) {
			latestMachine, latestErr := s.store.GetMachine(machineID)
			if latestErr == nil {
				data.Machine = latestMachine
			}
			data.ErrorMessage = "この洗濯機は利用開始できません。空きの状態を確認してから再度お試しください。"
			s.renderFragment(w, "machine_action_panel", data)
			return
		}

		http.Error(w, "failed to start session", http.StatusBadRequest)
		return
	}

	data.FlashMessage = "利用開始を受け付けました。通知が必要な場合は利用中画面から設定できます。"
	s.renderFragment(w, "machine_action_panel", data)
}

func (s *Server) handleClaimSession(w http.ResponseWriter, r *http.Request) {
	sessionID := chi.URLParam(r, "sessionID")
	err := s.store.ClaimSession(sessionID)

	data := pageData{
		CurrentSession:         s.store.CurrentSession(),
		NotificationPreference: s.store.NotificationPreference(),
		Now:                    s.store.Now(),
	}

	if err != nil {
		data.ErrorMessage = "回収完了を登録できませんでした。もう一度お試しください。"
		s.renderFragment(w, "current_session_panel", data)
		return
	}

	data.FlashMessage = "回収完了を受け付けました。"
	s.renderFragment(w, "current_session_panel", data)
}

func (s *Server) handleEnableNotifications(w http.ResponseWriter, r *http.Request) {
	session := s.store.EnableNotifications()

	data := pageData{
		CurrentSession:         session,
		NotificationPreference: s.store.NotificationPreference(),
		FlashMessage:           "通知を有効にしました。洗濯終了が近づいたらお知らせします。",
		Now:                    s.store.Now(),
	}

	if session != nil {
		s.renderFragment(w, "current_session_panel", data)
		return
	}

	s.renderFragment(w, "notification_panel", data)
}

func (s *Server) renderPage(w http.ResponseWriter, page string, data pageData) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	templates, err := parseTemplates(
		filepath.Join(s.root, "web", "templates", "layout", "*.gohtml"),
		filepath.Join(s.root, "web", "templates", "partials", "*.gohtml"),
		filepath.Join(s.root, "web", "templates", "pages", fmt.Sprintf("%s.gohtml", page)),
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := templates.ExecuteTemplate(w, fmt.Sprintf("%s_page", page), data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (s *Server) renderFragment(w http.ResponseWriter, name string, data pageData) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	templates, err := parseTemplates(filepath.Join(s.root, "web", "templates", "partials", "*.gohtml"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := templates.ExecuteTemplate(w, name, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
