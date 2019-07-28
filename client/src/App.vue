<template>
    <div class="row m-0 h-100" id="app">
        <div class="col-12 col-sm-2 col-md-4 d-none d-sm-block"></div>
        <Login @login="login" v-if="token == null"></Login>
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
                token: null
            }
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
