<template>
    <div class="row m-0 h-100" id="app">
        <div class="col-12 col-sm-2 col-md-4 d-none d-sm-block"></div>
        <Login :preset="presetLogin"
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
                presetLogin: {
                    host: "",
                    username: "",
                    password: ""
                },
            };
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
                    this.presetLogin = res;
                })
                .catch(err => console.error(err));
        },
        methods: {
            login(host, username, password) {
                let options = {
                    method: 'POST',
                    body: JSON.stringify({
                        hostname: host,
                        username: username,
                        password: password,
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
