async function apiGet<T>(url: string): Promise<T> {
    const res = await fetch(url)
    if (!res.ok) throw new Error(`${res.status} ${res.statusText}`)
    const { data } = await res.json()
    return data
}

export { apiGet }