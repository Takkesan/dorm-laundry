# dorm-laundry

寮の共同洗濯機向けフロントエンドです。`AGENTS.md` に合わせて、`Go + chi + html/template + htmx + Tailwind CSS` のサーバーサイドレンダリング構成にしています。

## 画面

- `/` 洗濯機一覧
- `/machines/{machineId}` 洗濯機詳細と利用開始導線
- `/sessions/current` 自分の利用中セッション

## 開発

```bash
npm install
npm run build:css
go run ./cmd/server
```

## テスト

```bash
go test ./...
```

## メモ

- `htmx` は `web/static/js/htmx.min.js` にローカル配置しています
- 洗濯機アイコンは `web/static/icons/washer.svg` を利用しています
