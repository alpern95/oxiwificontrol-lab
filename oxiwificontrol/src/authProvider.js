// in src/authProvider.js

const authProvider = {
    login: ({ username, password }) =>  {
        const request = new Request('http://192.168.1.32:3000/api/v1/user/login', {
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
            .then(({ token }) => {
                //const decodedToken = decodeJwt(token);
                localStorage.setItem('token', token);
                localStorage.setItem('username', username);
                return Promise.resolve();
                //localStorage.setItem('permissions', decodedToken.permissions);
            });


    },
    logout: () => {
        localStorage.removeItem('username');
        return Promise.resolve();
    },
    checkError: () => Promise.resolve(),
    checkAuth: () =>
        localStorage.getItem('username') ? Promise.resolve() : Promise.reject(),
    getPermissions: () => Promise.reject('Unknown method'),
};

export default authProvider
