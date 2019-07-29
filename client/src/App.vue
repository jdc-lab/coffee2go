<template>
    <div class="row m-0 h-100" id="app">
        <div class="col-12 col-sm-2 col-md-4 d-none d-sm-block"></div>
        <Login :host-preset="host"
               :password-preset="password"
               :username-preset="username"
               @login="login"
               v-if="token == null"/>
        <p v-if="token != null">Logged In</p>
        <div class="col-12 col-sm-2 col-md-4 d-none d-sm-block"></div>
    </div>
</template>

<script>
    import Login from './components/Login.vue'
    import '../public/css/default.css'

    export default {
        name: 'app',
        components: {
            Login
        },
        data: function () {
            return {
                token: null,
                host: "",
                username: "",
                password: ""
            }
        },
        created() {
            let options = {
                method: 'GET',
                headers: {
                    'Content-Type': 'application/json'
                }
            };

            fetch('/api/login/preset', options)
                .then(data => {
                    return data.json()
                })
                .then(res => {
                    if (res.host)
                        this.host = res.host;
                    if (res.username)
                        this.username = res.username;
                    if (res.password)
                        this.password = res.password;
                })
                .catch(err => console.error(err));
        },
        methods: {
            login() {
                let options = {
                    method: 'POST',
                    body: JSON.stringify({
                        hostname: this.host,
                        username: this.username,
                        password: this.password,
                    }),
                    headers: {
                        'Content-Type': 'application/json'
                    }
                };

                fetch('/api/login', options)
                    .then(data => {
                        return data.json()
                    })
                    .then(res => {
                        this.token = res.token
                    })
                    .catch(err => console.error(err));
            }
        }
    }
</script>
