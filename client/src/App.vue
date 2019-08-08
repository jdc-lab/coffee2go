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
    import axios from 'axios'

    export default {
        name: 'app',
        components: {
            Login
        },
        data: function () {
            return {
                sse: null,
                token: null,
                presetLogin: {
                    host: "",
                    username: "",
                    password: ""
                },
            };
        },
        created() {
            axios.get('/api/login/preset').then(res => {
                this.presetLogin = res.data;
            }).catch(err => console.error(err));
        },
        methods: {
            login(host, username, password) {
                axios.post('/api/login', {
                    host: host,
                    username: username,
                    password: password,
                }).then(res => {
                    this.token = res.data.token
                    this.subscribe()
                }).catch(err => console.error(err));
            },
            /**
             * Subscribe to push messages
             */
            subscribe() {

                axios.get('/api/auth/' + this.token + '/push/register').then(res => {
                    if (this.sse == null) {
                        this.sse = new EventSource("/push/" + res.data);
                        this.sse.addEventListener("recv", function (event) {
                            console.log("recv ")
                        });
                        this.sse.onopen = function (event) {
                            console.log("open ")
                        };
                        this.sse.onerror = function (event) {
                            console.log("error ")
                        };

                    }
                }).catch(err => console.error(err));

                /*

                this.sse.onmessage = function(e) {
                    console.log(e.channel)
                    console.log(e.data)
                    console.log(e.time)
                }*/
            }
        }
    }
</script>
