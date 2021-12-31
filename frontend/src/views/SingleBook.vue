<template>
  <v-container>
    <v-app id="single">
      <v-dialog v-model="dialog" width="500">
        <v-card>
          <v-card-title class="text-h5"> </v-card-title>

          <v-card-text class="text-center p-5">
            <h1>
              Are you sure you want to <span class="red--text">Delete</span>?
            </h1>
          </v-card-text>

          <v-divider></v-divider>

          <v-card-actions>
            <v-spacer></v-spacer>
            <v-btn depressed @click="dialog = false"> Cancel </v-btn>
            <v-btn depressed color="error" @click="deleteBook(deleteBook.id)">
              Delete
            </v-btn>
          </v-card-actions>
        </v-card>
      </v-dialog>
      <v-container grid-list-md>
        <v-layout row wrap class="justify-space-between">
          <v-container grid-list-md text-xs-center>
            <v-layout row wrap>
              <v-flex xs5>
                <v-card>
                  <v-img
                    id="product-img"
                    src="https://cdn.shopify.com/s/files/1/0054/3214/8081/products/45_665ca830-d1f2-48bd-8c84-02375dc2d883_large.jpg?v=1544128468"
                    height="100%"
                    class="grey darken-4"
                  ></v-img>
                </v-card>
              </v-flex>
              <v-flex xs7 pl-5 class="d-flex text-xs-left">
                <div class="product-summary">
                  <h2 class="product-title">{{ singleBook.name }}</h2>
                  <div class="user-ratings">
                    <div class="star-rating">
                      <span style="width: 100%"></span>
                    </div>
                  </div>
                  <div class="price">
                    <h3>
                      {{ singleBook.author }}
                    </h3>
                    <p>{{ singleBook.pageNumber }} pages</p>
                  </div>
                  <div class="prodect-details">
                    <h3>Product Details</h3>
                    <p>
                      {{ singleBook.description }}
                    </p>
                    <v-btn depressed color="error" @click="dialog = true">
                      Delete
                    </v-btn>
                  </div>
                </div>
              </v-flex>
            </v-layout>
          </v-container>
        </v-layout>
      </v-container>
    </v-app>
  </v-container>
</template>

<script>
import axios from "axios";
import axiosConfig from "../axios.config";
export default {
  name: "HelloWorld",
  data() {
    return {
      dialog: false,
      singleBook: {},
    };
  },
  mounted() {
    this.getSingleBook();
  },
  methods: {
    getSingleBook() {
      axios
        .get(`/books/${this.$route.query.id}`, axiosConfig)
        .then((response) => (this.books = response.data.data));
    },
    deleteBook(id) {
      axios
        .delete(`/books/${id}`, axiosConfig)
        .then(() => this.$router.push("/"));
    },
  },
};
</script>
