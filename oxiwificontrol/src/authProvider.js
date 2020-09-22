// in src/authProvider.js

const authProvider = {
    login: ({ username, password }) =>  {
        const request = new Request('https://192.168.112.11:4431/api/v1/user/login', {
            method: 'POST',
            body: JSON.stringify({ username, password }),
            headers: new Headers({ 'Content-Type': 'application/json' }),
        });
        return fetch(request)
            .then(response => {
                if (response.status < 200 || response.status >= 300) {
                    throw new Error(response.statusText);
                }
                return response.json();
            })
            .then(({ token,role}) => {
                localStorage.setItem('token', token);
                localStorage.setItem('username', username);
                localStorage.setItem('permissions', role);
                return Promise.resolve();
            })
    },

    logout: () => {
        localStorage.removeItem('username');
        localStorage.removeItem('token');
        localStorage.removeItem('permissions');
        return Promise.resolve();
    },

    checkError: () => Promise.resolve(),

    checkAuth: () => {
        //localStorage.getItem('username') ? Promise.resolve() : Promise.reject(),
        return localStorage.getItem('token') ? Promise.resolve() : Promise.reject();
    },

    getPermissions: () => {
        const role = localStorage.getItem('permissions');
        return role ? Promise.resolve(role) : Promise.reject();
    },
};
export default authProvider
