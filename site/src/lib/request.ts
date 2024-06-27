import { env } from "$env/dynamic/public"

export const request = async(path: string, method: string, body?: any) => {
    let h = new Headers()
    body = body ?? undefined
    h.append('content-type', 'application/json')
    const r = await fetch(env.PUBLIC_API + path, {
        method: method,
        body: JSON.stringify(body),
        headers: h
    })
    return await r.json()
}