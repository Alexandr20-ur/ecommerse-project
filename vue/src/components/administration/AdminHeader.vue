<script setup>

</script>

<template>
  <header class="d-flex p-3 lg:px-5 align-items-center border-bottom h-14">

    <button class="btn btn-default btn-icon mr-3 lg:d-none" onclick="open_dialog('dialog_menu_mobile')">
      <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" width="24" height="24" color="currentColor" fill="none">
        <path d="M4 5L20 5" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" />
        <path d="M4 12L20 12" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" />
        <path d="M4 19L20 19" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" />
      </svg>
    </button>

    <h5 class="m-0 text-xl max-sm:d-none">Главная </h5>
    <div class="d-flex align-items-center gap-2 ml-auto">
      <form>
        <input type="search" class="form-control" list="search_suggest" placeholder="Поиск">
      </form>
      <button class="btn btn-neutral btn-icon">
        <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="lucide lucide-bell-icon lucide-bell"><path d="M10.268 21a2 2 0 0 0 3.464 0"/><path d="M3.262 15.326A1 1 0 0 0 4 17h16a1 1 0 0 0 .74-1.673C19.41 13.956 18 12.499 18 8A6 6 0 0 0 6 8c0 4.499-1.411 5.956-2.738 7.326"/></svg>
      </button>
      <a class="btn btn-neutral" href="#">
        <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" width="20" height="20" fill="#000" color="var(--neutral-600)">
          <path d="M17.8063 14.8372C17.9226 14.9064 18.0663 14.9875 18.229 15.0793C18.9418 15.4814 20.0193 16.0893 20.7575 16.8118C21.2191 17.2637 21.6578 17.8592 21.7375 18.5888C21.8223 19.3646 21.4839 20.0927 20.8048 20.7396C19.6334 21.8556 18.2276 22.75 16.4093 22.75H7.59104C5.77274 22.75 4.36695 21.8556 3.1955 20.7396C2.51649 20.0927 2.17802 19.3646 2.26283 18.5888C2.34257 17.8592 2.78123 17.2637 3.2429 16.8118C3.98106 16.0893 5.05857 15.4814 5.77139 15.0793C5.93404 14.9875 6.07773 14.9064 6.19404 14.8372C9.74809 12.7209 14.2523 12.7209 17.8063 14.8372Z" fill="currentColor"></path>
          <path d="M6.75018 6.5C6.75018 3.6005 9.10068 1.25 12.0002 1.25C14.8997 1.25 17.2502 3.6005 17.2502 6.5C17.2502 9.39949 14.8997 11.75 12.0002 11.75C9.10068 11.75 6.75018 9.39949 6.75018 6.5Z" fill="currentColor"></path>
        </svg>
        Мой акаунт
      </a>

    </div>
  </header>
</template>

<style scoped>

</style>