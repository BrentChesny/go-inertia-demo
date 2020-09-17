import React from 'react';
import { InertiaLink, usePage } from '@inertiajs/inertia-react';

const Second = () => {
    // const page = usePage();
    // console.log(page);
    return (
            <div>
                <h1>This is the second page.</h1>
                <InertiaLink href={"/"}>
                    Back to Home
                </InertiaLink>
            </div>
    );
};

export default Second