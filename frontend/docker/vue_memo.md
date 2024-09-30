<!-- Ctrl+K, V -->

# vue プロジェクト作成

```shell
npm create vue@latest
```

CLI で進むので下記質問で Yes

- AddTypeScript?
- Add Vue Router for Single Page Application developement?
- AddESLint for code quality?
- Add Prettier for code formatting?

下記コマンドで初期化

```shell
npm install
npm run format
```

Vue Router を No にした場合は下記コマンドで Router を追加できる

```shell
npm i vue-router@4
```

index.html の lang を en から ja に変更しておく

```html
<html lang="ja"></html>
```

docker からだと繋がらないので vite.config.js を修正。host: true を追加する。
ついでに hotreload の設定も(watch)。

```js
import { defineConfig } from 'vite';
import vue from '@vitejs/plugin-vue';

// https://vitejs.dev/config/
export default defineConfig({
  //ここを追加！
  server: {
    host: true,
    // hotreload の設定
    watch: {
      usePolling: true,
    },
  },

  plugins: [vue()],
});
```

```
npm run dev
```

http://localhost:5173/
にアクセスして動作確認

## Vuetify インストール

マテリアルデザインのフレームワーク  
https://vuetifyjs.com/ja/getting-started/installation/#section-30a430f330b930c830fc30eb

使い方はこちらが分かりやすい  
https://zenn.dev/bbled/books/vuetify3_book/viewer/sec1_2

既存プロジェクトで使う場合、下記を実行

```shell
npm i vuetify
npm i @mdi/font -D
```

main.ts にコードを追加

```ts
import { createApp } from 'vue';

// Vuetify
import 'vuetify/styles';
import { createVuetify } from 'vuetify';
import * as components from 'vuetify/components';
import * as directives from 'vuetify/directives';
import '@mdi/font/css/materialdesignicons.css'; // マテリアルデザインアイコン

// Components
import App from './App.vue';

const vuetify = createVuetify({
  components,
  directives,
  theme: {
    defaultTheme: 'dark',
  },
});

createApp(App).use(vuetify).mount('#app');
```

## アプリの骨組み作成

App.vue を下記のような感じで

```html
<script setup lang="ts">
  import { ref } from 'vue';

  const drawer = ref(false);
</script>

<template>
  <v-app>
    <v-app-bar color="primary">
      <template v-slot:prepend>
        <v-app-bar-nav-icon
          variant="text"
          @click.stop="drawer = !drawer"
        ></v-app-bar-nav-icon>
      </template>

      <v-app-bar-title> タイトル </v-app-bar-title>

      <template v-slot:append>
        <v-btn>
          <v-icon> mdi-home </v-icon>
        </v-btn>
      </template>
    </v-app-bar>

    <v-navigation-drawer v-model="drawer">
      <v-list>
        <v-list-item
          prepend-icon="mdi-view-dashboard"
          title="Home"
          value="home"
        ></v-list-item>
        <v-list-item
          prepend-icon="mdi-forum"
          title="About"
          value="about"
        ></v-list-item>
      </v-list>
    </v-navigation-drawer>

    <v-main>
      <v-container>
        <v-sheet> メインコンテンツ </v-sheet>
      </v-container>
    </v-main>
  </v-app>
</template>

<style scoped></style>
```

## Router 設定

src 配下に router フォルダを作成し、index.ts を作成

```ts
import { createRouter, createWebHistory } from 'vue-router';

const router = createRouter({
  history: createWebHistory(),
  routes: [],
});

export default router;
```

main.ts にコードを追加

```ts
import './assets/main.css'

import { createApp } from 'vue'
import App from './App.vue'
import router from './router' // 追加
：
// .use(router) 追加
createApp(App).use(vuetify).use(router).mount('#app')
```

App.vue のページ遷移部分を &lt;router-view /&gt; に変更

```html
<v-main>
  <v-container>
    <v-sheet>
      <router-view />
    </v-sheet>
  </v-container>
</v-main>
```

router/index.ts でルート設定追加

```ts
const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      path: '/',
      name: 'mainTodo',
      component: MainTodo,
    },
  ],
});
```

クリックで画面遷移するには $router.push() を使用

```html
<v-navigation-drawer v-model="drawer">
  <v-list>
    <v-list-item
      prepend-icon="mdi-view-dashboard"
      title="Home"
      value="home"
      @click="$router.push('/')"
    />
    <v-list-item
      prepend-icon="mdi-forum"
      title="About"
      value="about"
      @click="$router.push('/about')"
    />
  </v-list>
</v-navigation-drawer>
```

router の component 指定で遅延ローディングを設定可能

```ts
{
  path: '/about',
  name: 'about',
  component: () => import('@/components/TheAbout.vue')
}
```

404 Not found 追加

```ts
{
  path: '/:pathMatch(.*)*',
  name: 'notFound',
  component: () => import('@/components/NotFound.vue')
}
```

パラメータ

```ts
  path: '/about/:id',
  name: 'about',
  component: () => import('@/components/TheAbout.vue')
```

```html
<script setup lang="ts">
  const route = useRoute();
  const id = ref(route.params.id);
</script>
```

## TIps

### Floating Button

```ts
<v-layout-item model-value position="bottom" class="text-center" size="88">
  <div class="ma-4">
    <v-btn icon="mdi-plus" size="large" color="primary" elevation="8" />
  </div>
</v-layout-item>
```
