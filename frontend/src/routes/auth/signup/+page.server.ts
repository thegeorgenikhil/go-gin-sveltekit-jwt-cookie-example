import { type Actions, error, redirect } from "@sveltejs/kit";
import type { PageServerLoad } from "./$types";

export const load: PageServerLoad = async ({ locals }) => {
    // redirect user if logged in
    if (locals.token) {
        throw redirect(302, '/')
    }
}

export const actions: Actions = {
    signup: async ({ fetch, request, cookies }) => {
        const form = await request.formData();
        const email = form.get('email');
        const password = form.get('password');

        const res = await fetch("http://localhost:8080/auth/signup", {
            method: "POST",
            body: JSON.stringify({ email, password }),
        })

        const data = await res.json();

        if (res.status === 201) {
            cookies.set('token', data.token, {
                // send cookie for every page
                path: '/',
                // server side only cookie so you can't use `document.cookie`
                httpOnly: true,
                // only requests from same site can send cookies
                // https://developer.mozilla.org/en-US/docs/Glossary/CSRF
                sameSite: 'strict',
                // only sent over HTTPS in production
                secure: true,
                // set cookie to expire after a month
                maxAge: 60 * 60 * 24 * 30,
            })
            throw redirect(302, "/user/id");
        } else {
            return error(422, {
                message: data.message
            });
        }
    },
};

