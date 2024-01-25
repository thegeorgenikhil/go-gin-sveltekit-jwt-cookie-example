import type { Handle } from '@sveltejs/kit'

export const handle: Handle = async ({ event, resolve }) => {
    // get cookies from browser
    const token = event.cookies.get('token')

    if (!token) {
        // if there is no token load page as normal
        return await resolve(event)
    }

    // if `token` exists set `events.local`
    if (token) {
        event.locals.token = token
    }

    // load page as normal
    return await resolve(event)
}
