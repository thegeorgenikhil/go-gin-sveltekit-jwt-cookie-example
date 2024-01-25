import type { LayoutServerLoad } from './$types'

// get `locals.token` and pass it to the `page` store
export const load: LayoutServerLoad = async ({ locals }) => {
    return {
        token: locals.token,
    }
}
