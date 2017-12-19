<template>
  <v-container fluid>
    <v-layout row>
      <v-flex xs12 sm8 offset-sm2>
        <p class="headline">Sign In</p>
        <v-form v-model="valid" ref="form" lazy-validation>
          <v-text-field
            label="Username"
            v-model="username"
            :rules="usernameRules"
            required
          ></v-text-field>
          <v-text-field
            label="Password"
            v-model="password"
            :rules="passwordRules"
            required
            type="password"
          ></v-text-field>
          <v-btn
            @click="submit"
            :disabled="!valid"
          >
            Submit
          </v-btn>
        </v-form>
        <v-progress-linear :indeterminate="true" v-show="sending"></v-progress-linear>
        <v-alert color="error" v-show="errors.length > 0" icon="warning" value="true">
          <p v-for="(error, i) in errors" :key="i">
            {{ error.message }}
          </p>
        </v-alert>
      </v-flex>
    </v-layout>
  </v-container>
</template>

<script>
import { HTTP } from '../../api'

export default {
  data: () => ({
    sending: false,
    username: '',
    usernameRules: [v => !!v || 'Username is required'],
    password: '',
    passwordRules: [v => !!v || 'Password is required'],
    errors: [],
    valid: true
  }),
  methods: {
    submit() {
      if (this.$refs.form.validate()) {
        this.sending = true
        HTTP.post('/auth', {
          username: this.username,
          password: this.password
        })
          .then(response => {
            this.sending = false
            let token = response.data.token.toString()
            localStorage.setItem('token', token)
            this.$store.dispatch('setIsLoggedIn', true)
            this.$router.push({ name: 'Hello' })
          })
          .catch(e => {
            this.sending = false
            this.errors.push(e)
          })
      }
    }
  }
}
</script>

<style>

</style>
