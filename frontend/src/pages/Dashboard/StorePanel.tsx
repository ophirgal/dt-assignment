import { useParams } from 'react-router-dom'
import { STORES } from '@/data/stores'

export default function StorePanel() {
    const { storeSystemName } = useParams()
    const store = STORES.find((s) => s.systemName === storeSystemName)

    return (
        <main className="p-8">
            <h1 className="text-2xl font-medium text-foreground">
                {store?.displayName ?? 'Store Not Found'}
            </h1>
        </main>
    )
}
