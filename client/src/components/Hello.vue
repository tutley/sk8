<template>
  <v-container fluid>
  <v-layout>
      <v-flex d-flex>
        <v-card>
          <v-card-media
            :height="pictureHeight"
            class="white--text"
            src="/static/img/skate.webp">
            <v-layout column align-center justify-center>
            <span class="display-4 white--text">{{msg}}</span>
            </v-layout>
          </v-card-media>
        </v-card>
      </v-flex>
  </v-layout>
    <v-layout>
      <v-flex>
          <div class="hello">
            <span class="subheading">
              Hello, welcome to sk8! This is a place where you can post about skating! Just login and start posting examples for other skaters!
              </span>
             <center>
              <v-btn color="primary" to="/examples">Get Started</v-btn>
              </center>
          </div>
      </v-flex>
    </v-layout>
  </v-container>
</template>

<script>
export default {
  name: 'hello',
  data() {
    return {
      msg: 'We Skate'
    }
  },
  computed: {
    pictureHeight() {
      return this.$store.getters.viewHeight
    }
  }
}
</script>

<style>
.container {
  padding: 0px;
}
.hello {
  padding: 12px;
}
</style>
