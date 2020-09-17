import React from 'react';
import { InertiaLink, usePage } from '@inertiajs/inertia-react';

export default () => {
    const page = usePage();
    const message = page.name ? `Welcome, ${page.name}!` : 'Welcome!'
    return (
            <div>
                <h1>{message}</h1>
                <InertiaLink href={"/second"}>
                    A link to another page
                </InertiaLink>
            </div>
    );
};
