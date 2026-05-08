import { useMemo, useState } from 'react'
import { useParams } from 'react-router-dom'
import { useStores } from '@/hooks/use-stores'
import { useForecast } from '@/hooks/use-forecast'
import type { Forecast } from '@/models/forecast'

// ── types ────────────────────────────────────────────────────────────────────

interface Product { id: string; name: string }
interface HourRow { hour: number; perProduct: Record<string, number>; total: number }

// ── constants ─────────────────────────────────────────────────────────────────

const CHART_HEIGHT = 340
const PALETTE = [
    'oklch(0.55 0.13 35)',  // terracotta
    'oklch(0.70 0.12 65)',  // amber
    'oklch(0.78 0.09 90)',  // sand
    'oklch(0.45 0.08 50)',  // umber
    'oklch(0.62 0.06 40)',  // muted clay
]

// ── helpers ───────────────────────────────────────────────────────────────────

// Using January 10th 2026 harcoded for assignment. In a real app, we would use the current day.
function january10_2026() {
    return new Date('2026-01-10').toISOString().slice(0, 10)
}

function fmtDate(iso: string) {
    return new Date(iso + 'T00:00:00').toLocaleDateString('en-US', {
        weekday: 'long', month: 'long', day: 'numeric',
    })
}

const pad = (h: number) => String(h).padStart(2, '0')

// ── StorePanel ────────────────────────────────────────────────────────────────

export default function StorePanel() {
    const { storeSystemName } = useParams<{ storeSystemName: string }>()
    const [date, setDate] = useState(january10_2026())

    const { data: stores = [] } = useStores()
    const storeId = stores.find((s) => s.systemName === storeSystemName)?.id
    const { data: forecasts = [] } = useForecast(storeId!, date)

    const store = stores.find((s) => s.systemName === storeSystemName)

    const products = useMemo<Product[]>(() =>
        [...new Set(forecasts?.map((f) => f.productName))].map((name) => ({ id: name, name })),
        [forecasts],
    )

    const rows = useMemo<HourRow[]>(() =>
        Array.from({ length: 24 }, (_, hour) => {
            const perProduct: Record<string, number> = {}
            forecasts
                ?.filter((f: Forecast) => f.hour === hour)
                .forEach((f: Forecast) => { perProduct[f.productName] = Math.ceil(f.predictedQuantity) })
            const total = Object.values(perProduct).reduce((s, v) => s + v, 0)
            return { hour, perProduct, total }
        }),
        [forecasts],
    )

    return (
        <div className="flex flex-col h-full">
            <ForecastHeader
                storeName={store?.displayName ?? '…'}
                date={date}
                onDateChange={setDate}
            />
            <div className="flex-1 overflow-auto">
                <ForecastChart rows={rows} products={products} />
            </div>
            <ForecastFooter />
        </div>
    )
}


// ── ForecastHeader ────────────────────────────────────────────────────────────

interface HeaderProps {
    storeName: string
    date: string
    onDateChange: (d: string) => void
}

function ForecastHeader({ storeName, date, onDateChange }: HeaderProps) {
    return (
        <header className="flex items-end justify-between px-10 pt-8 pb-6 border-b border-border">
            <div>
                <p className="text-[11px] tracking-[0.14em] uppercase text-muted-foreground font-medium mb-2">
                    Forecast · {storeName}
                </p>
                <h1 className="text-4xl font-medium tracking-tight leading-none">
                    {fmtDate(date)}
                </h1>
            </div>
            <div className="flex flex-col items-end gap-2">
                <label htmlFor="forecast-date" className="text-[11px] tracking-[0.14em] uppercase text-muted-foreground font-medium">
                    Forecast Date
                </label>
                <input
                    id="forecast-date"
                    type="date"
                    value={date}
                    onChange={(e) => onDateChange(e.target.value)}
                    className="font-mono text-sm bg-background border border-border rounded px-3 py-2 outline-none focus:border-primary cursor-pointer transition-colors"
                />
            </div>
        </header>
    )
}

// ── HoverDetail ───────────────────────────────────────────────────────────────

function HoverDetail({ row, products }: { row: HourRow; products: Product[] }) {
    return (
        <div className="mt-6 px-5 py-4 border border-border rounded-lg flex items-center gap-8 font-mono">
            <div className="shrink-0">
                <p className="text-[11px] tracking-[0.12em] uppercase text-muted-foreground mb-1">
                    {pad(row.hour)}:00 – {pad(row.hour + 1)}:00
                </p>
                <p className="text-3xl font-medium leading-none">
                    {row.total}<span className="text-sm text-muted-foreground ml-2">items</span>
                </p>
            </div>
            <div className="grid grid-cols-5 gap-6 flex-1">
                {products.map((p, i) => (
                    <div key={p.id}>
                        <div className="flex items-center gap-1.5 mb-1">
                            <span className="w-2 h-2 rounded-sm shrink-0" style={{ background: PALETTE[i] }} />
                            <span className="text-[11px] text-muted-foreground truncate">{p.name}</span>
                        </div>
                        <p className="text-lg font-medium">{row.perProduct[p.id] ?? 0}</p>
                    </div>
                ))}
            </div>
        </div>
    )
}

// ── ForecastChart ─────────────────────────────────────────────────────────────

function ForecastChart({ rows, products }: { rows: HourRow[]; products: Product[] }) {
    const [hoverHour, setHoverHour] = useState<number | null>(null)
    const maxTotal = Math.max(...rows.map((r) => r.total), 1)
    const ticks = [0, 0.33, 0.66, 1]

    return (
        <div className="px-10 pt-6 pb-4">
            {/* Legend */}
            <div className="flex gap-5 mb-6 flex-wrap">
                {products.map((p, i) => (
                    <div key={p.id} className="flex items-center gap-2 text-[13px] text-muted-foreground">
                        <span className="w-2.5 h-2.5 rounded-sm" style={{ background: PALETTE[i] }} />
                        {p.name}
                    </div>
                ))}
            </div>

            {/* Chart */}
            <div className="relative">
                {/* Y-axis */}
                <div className="absolute left-0 top-0 flex items-center gap-1" style={{ height: CHART_HEIGHT }}>
                    <span
                        className="font-mono text-[11px] tracking-[0.14em] uppercase text-muted-foreground"
                        style={{ writingMode: 'vertical-rl', transform: 'rotate(180deg)' }}
                    >
                        Items
                    </span>
                    <div
                        className="flex flex-col justify-between text-right pr-3 font-mono text-[11px] text-muted-foreground w-8"
                        style={{ height: CHART_HEIGHT }}
                    >
                        {[...ticks].reverse().map((t) => (
                            <span key={t}>{Math.round(maxTotal * t)}</span>
                        ))}
                    </div>
                </div>

                <div className="ml-14">
                    {/* Bars + grid */}
                    <div className="relative" style={{ height: CHART_HEIGHT }}>
                        {ticks.map((t) => (
                            <div
                                key={t}
                                className="absolute left-0 right-0 border-t border-border/60"
                                style={{ top: `${(1 - t) * 100}%` }}
                            />
                        ))}
                        <div
                            className="absolute inset-0 grid items-end"
                            style={{ gridTemplateColumns: 'repeat(24, 1fr)', gap: 5 }}
                        >
                            {rows.map((row) => {
                                const barH = (row.total / maxTotal) * CHART_HEIGHT
                                const isHovered = hoverHour === row.hour
                                return (
                                    <div
                                        key={row.hour}
                                        className="relative flex flex-col justify-end"
                                        style={{ height: CHART_HEIGHT }}
                                        onMouseEnter={() => setHoverHour(row.hour)}
                                        onMouseLeave={() => setHoverHour(null)}
                                    >
                                        {isHovered && row.total > 0 && (
                                            <div
                                                className="absolute left-1/2 -translate-x-1/2 font-mono text-xs font-semibold pointer-events-none"
                                                style={{ bottom: barH + 6 }}
                                            >
                                                {row.total}
                                            </div>
                                        )}
                                        <div
                                            className="w-full flex flex-col-reverse rounded-t-sm overflow-hidden transition-opacity duration-150"
                                            style={{
                                                height: barH,
                                                opacity: hoverHour !== null && !isHovered ? 0.4 : 1,
                                            }}
                                        >
                                            {products.map((p, i) => {
                                                const qty = row.perProduct[p.id] ?? 0
                                                if (qty === 0 || row.total === 0) return null
                                                return (
                                                    <div
                                                        key={p.id}
                                                        style={{ height: `${(qty / row.total) * 100}%`, background: PALETTE[i] }}
                                                    />
                                                )
                                            })}
                                        </div>
                                    </div>
                                )
                            })}
                        </div>
                    </div>

                    {/* X-axis */}
                    <div
                        className="grid mt-2 font-mono text-[11px] text-muted-foreground text-center"
                        style={{ gridTemplateColumns: 'repeat(24, 1fr)', gap: 5 }}
                    >
                        {rows.map((row) => (
                            <div
                                key={row.hour}
                                className={row.hour % 2 !== 0 && hoverHour !== row.hour ? 'opacity-0' : ''}
                            >
                                {pad(row.hour)}
                            </div>
                        ))}
                    </div>
                    <p className="mt-1 text-center font-mono text-[11px] tracking-[0.14em] uppercase text-muted-foreground">
                        Hour
                    </p>

                    {hoverHour !== null && (rows[hoverHour]?.total ?? 0) > 0 && (
                        <HoverDetail row={rows[hoverHour]!} products={products} />
                    )}
                </div>
            </div>
        </div>
    )
}

// ── ForecastFooter ────────────────────────────────────────────────────────────

function ForecastFooter() {
    return (
        <footer className="px-10 py-4 border-t border-border flex flex-wrap gap-5 font-mono text-xs text-muted-foreground">
            <span>predicted_quantity = ⌈avg(last 7 days, same hour)⌉</span>
            <span>·</span>
            <span>store-local time</span>
        </footer>
    )
}