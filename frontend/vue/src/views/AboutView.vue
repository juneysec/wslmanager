<script setup lang="ts">
import { ref } from 'vue'
import nodejslicenses from '../assets/node-js-licenses.json'
import golicenses from '../assets/go-licenses.json'

const tab = ref(0)
</script>

<template>
  <v-sheet>
    <h1>WSL Manager Ver1.0.0</h1>

    <v-container class="text-left">
      Vue の練習に作ったアプリ。以下のパッケージを利用して作成されています。
    </v-container>
    <v-tabs v-model="tab">
      <v-tab value="0">Node.js</v-tab>
      <v-tab value="1">Go</v-tab>
    </v-tabs>

    <v-window v-model="tab">
      <v-window-item value="0">
        <v-sheet>
          <v-table>
            <thead>
              <tr>
                <th>パッケージ</th>
                <th>ライセンス</th>
                <th>Publisher</th>
                <th>リポジトリ</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="(item, key) in nodejslicenses" :key="key">
                <td>{{ key }}</td>
                <td>{{ item.licenses }}</td>
                <td>{{ 'publisher' in item ? item.publisher : '' }}</td>
                <td>
                  <a :href="'repository' in item ? item.repository : ''">{{
                    'repository' in item ? item.repository : ''
                  }}</a>
                </td>
              </tr>
            </tbody>
          </v-table>
        </v-sheet>
      </v-window-item>

      <v-window-item value="1">
        <v-sheet>
          <v-table>
            <thead>
              <tr>
                <th>パッケージ</th>
                <th>ラインセス</th>
                <th>ライセンスURL</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="(item, key) in golicenses" :key="key">
                <td>{{ key }}</td>
                <td>{{ item.LicenseName }}</td>
                <td>
                  <a target="_blank" :href="item.LicenseURL">{{ item.LicenseURL }}</a>
                </td>
              </tr>
            </tbody>
          </v-table>
        </v-sheet>
      </v-window-item>
    </v-window>
  </v-sheet>
</template>

<style></style>
