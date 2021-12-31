<template>
  <v-container>
    <v-app id="inspire">
      <v-dialog v-model="newItem" width="500">
        <v-card>
          <v-card-title class="text-h5"> Create new item </v-card-title>

          <v-card-text class="text-center p-5">
            <v-form ref="form" v-model="valid" lazy-validation>
              <v-text-field
                v-model="data.name"
                :counter="20"
                :rules="nameRules"
                label="Name"
                required
              ></v-text-field>

              <v-text-field
                v-model="data.author"
                :rules="authorRules"
                label="Author"
                required
              ></v-text-field>

              <v-text-field
                v-model="data.printYear"
                :rules="printYearRules"
                label="Print year"
              ></v-text-field>

              <v-text-field
                v-model="data.pageNumber"
                label="Page number"
              ></v-text-field>

              <v-textarea
                name="input-6-1"
                label="Description"
                v-model="data.description"
              ></v-textarea>
            </v-form>
          </v-card-text>
          <v-card-action>
            <v-btn
              :disabled="!valid"
              color="success"
              class="mr-4"
              @click="validate"
            >
              Save
            </v-btn>

            <v-btn @click="newItem = false"> Cancel </v-btn>
          </v-card-action>
        </v-card>
      </v-dialog>
      <v-container grid-list-md>
        <v-btn @click="newItem = true">
          <span class="mr-2">Create New Book</span>
        </v-btn>
        <v-layout row wrap class="justify-space-between">
          <v-card
            class="my-6"
            max-width="374"
            v-for="item in books"
            :key="item.id"
          >
            <v-img height="250" :src="item.image"></v-img>

            <v-card-title>{{ item.name }}</v-card-title>

            <v-card-text>
              <div class="mb-4 text-subtitle-1">
                {{ item.author }} <span>{{ item.pageNumber }} pages</span>
              </div>
              <div class="mb-4 text-subtitle-1">{{ item.printYear }}</div>

              <div>
                {{ item.description }}
              </div>
            </v-card-text>

            <v-card-actions>
              <v-btn depressed color="primary" @click="pushToSingleBook(id)">
                View Book
              </v-btn>
            </v-card-actions>
          </v-card>
        </v-layout>
        <div class="text-center">
          <v-pagination :length="3" v-model="links.current_page" @input="getAllBooks"/>
        </div>
      </v-container>
    </v-app>
  </v-container>
</template>

<script>
import axios from "axios";
import axiosConfig from "../axios.config";
export default {
  name: "Home",
  data: () => ({
    newItem: false,
    books: {},
    valid: true,
    data: {
      name: "",
      author: "",
      printYear: "",
      pageNumber: 0,
      description: "",
    },
    links: {},
    nameRules: [
      (v) => !!v || "Name is required",
      (v) => (v && v.length <= 20) || "Name must be less than 20 characters",
    ],
    authorRules: [
      (v) => !!v || "Author is required",
      (v) => (v && v.length <= 20) || "Name must be less than 20 characters",
    ],
    printYearRules: [
      (v) => !!v || "Print year is required",
      (v) => (v && v.length <= 10) || "Name must be less than 10 characters",
    ],
  }),
  mounted() {
    this.getAllBooks(1);
  },
  methods: {
    validate() {
      this.$refs.form.validate();
      this.createNewBook();
    },
    reset() {
      this.$refs.form.reset();
    },
    resetValidation() {
      this.$refs.form.resetValidation();
    },
    pushToSingleBook(id) {
      this.$router.push({ name: "SingleBook", query: { id: id } });
    },
    getAllBooks(page) {
      axios.get(`/books?page=${page}`, axiosConfig).then((response) => {
        this.books = response.data.data;
        this.links = response.data.meta;
      });
    },
    createNewBook() {
      axios.post("/books", this.data, axiosConfig).then(() => {
        this.newItem = false;
        this.data = {
          name: "",
          author: "",
          printYear: "",
          pageNumber: 0,
          description: "",
        };
        this.getAllBooks(1);
      });
    },
  },
};
</script>
