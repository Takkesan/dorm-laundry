/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./web/templates/**/*.gohtml"],
  theme: {
    extend: {
      colors: {
        ink: "#2c2f31",
        line: "#d9dde0",
        outline: "#747779",
        accent: "#0058ba",
        "surface-base": "#f5f7f9",
        "surface-container-low": "#eef1f3",
        "surface-container-lowest": "#ffffff",
        "surface-container-high": "#dfe3e6",
        "surface-container-highest": "#d9dde0",
        "status-available": "#005f50",
        "status-available-soft": "#5bfedd",
        "status-running": "#005776",
        "status-running-soft": "#9adaff",
        "status-pickup": "#653e00",
        "status-pickup-soft": "#ffddb8",
        "status-available-panel": "#f2fffb",
        "status-running-panel": "#f1f9ff",
        "status-pickup-panel": "#fff7ef"
      },
      fontFamily: {
        headline: ["'Plus Jakarta Sans'", "sans-serif"],
        body: ["Manrope", "sans-serif"]
      },
      boxShadow: {
        card: "0 18px 40px rgba(44, 47, 49, 0.06)",
        float: "0 12px 30px rgba(0, 88, 186, 0.12)"
      },
      backgroundImage: {
        "primary-grad": "linear-gradient(135deg, #0058ba 0%, #6c9fff 100%)"
      },
      opacity: {
        15: "0.15"
      }
    }
  },
  plugins: []
};
