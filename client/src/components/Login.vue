<template>
    <div class="col-12 col-sm-8 col-md-4">
        <div class="my-5 text-center">
            <img alt="Coffee2Go" src="img/coffee2go.png"/>
        </div>
        <div autocomplete="off" class="my-5" id="form-login" spellcheck="false">
            <div class="form-group">
                <input :value="host" class="form-control border-none my-2 text-center" placeholder="Host" type="text">
                <input :value="username" class="form-control border-none my-2 text-center" placeholder="Username"
                       type="text">
                <input :value="password" class="form-control border-none my-2 text-center" placeholder="Password"
                       type="password">
                <div class="my-5 text-center">
                    <button @click="login" class="btn btn-primary px-3">
                        <small class="text-500">START CHATTING</small>
                    </button>
                </div>
            </div>
        </div>
        <p class="text-center">
            <small class="text-muted">&copy; 2019 Coffee2Go</small>
        </p>
    </div>
</template>

<script>
    export default {
        name: 'Login',
        props: {
            hostPreset: String,
            usernamePreset: String,
            passwordPreset: String
        },
        data: function () {
            return {
                host: this.hostPreset,
                username: this.usernamePreset,
                password: this.passwordPreset
            }
        },
        methods: {
            login() {
                this.$emit('login');
            }
        },
        watch: {
            hostPreset: function (newVal, oldVal) {
                if (oldVal == "")
                    this.host = newVal;
            },
            usernamePreset: function (newVal, oldVal) {
                if (oldVal == "")
                    this.username = newVal;
            },
            passwordPreset: function (newVal, oldVal) {
                if (oldVal == "")
                    this.password = newVal;
            }
        }
    }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
