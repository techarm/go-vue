<template>
    <aside class="sidenav navbar navbar-vertical navbar-expand-xs border-0 border-radius-xl my-3 fixed-start ms-3 bg-gradient-dark" id="sidenav-main">
        <div class="sidenav-header">
            <i class="fas fa-times p-3 cursor-pointer text-white opacity-5 position-absolute end-0 top-0 d-none d-xl-none" aria-hidden="true" id="iconSidenav"></i>
            <router-link class="navbar-brand m-0" to="/admin" target="_blank">
                <img src="/img/logo-ct.png" class="navbar-brand-img h-100" alt="main_logo">
                <span class="ms-1 font-weight-bold text-white">Teacharm Dashboard</span>
            </router-link>
        </div>
        <hr class="horizontal light mt-0 mb-2">
        <div class="collapse navbar-collapse w-auto" id="sidenav-collapse-main">
            <ul class="navbar-nav">
                <li class="nav-item">
                    <router-link class="nav-link text-white active bg-gradient-primary" to="/admin">
                        <div class="text-white text-center me-2 d-flex align-items-center justify-content-center">
                            <i class="material-icons opacity-10">dashboard</i>
                        </div>
                        <span class="nav-link-text ms-1">Dashboard</span>
                    </router-link>
                </li>
                <li class="nav-item mt-3">
                    <h6 class="ps-4 ms-2 text-uppercase text-xs text-white font-weight-bolder opacity-8">書籍管理</h6>
                </li>
                <li class="nav-item">
                    <router-link class="nav-link text-white" to="/admin/books/0">
                        <div class="text-white text-center me-2 d-flex align-items-center justify-content-center">
                            <i class="material-icons opacity-10">bookmark_add</i>
                        </div>
                        <span class="nav-link-text ms-1">書籍登録</span>
                    </router-link>
                </li>
                <li class="nav-item">
                    <router-link class="nav-link text-white" to="/admin/books">
                        <div class="text-white text-center me-2 d-flex align-items-center justify-content-center">
                            <i class="material-icons opacity-10">library_books</i>
                        </div>
                        <span class="nav-link-text ms-1">書籍一覧</span>
                    </router-link>
                </li>
                <li class="nav-item mt-3">
                    <h6 class="ps-4 ms-2 text-uppercase text-xs text-white font-weight-bolder opacity-8">ユーザー管理</h6>
                </li>
                <li class="nav-item">
                    <router-link class="nav-link text-white" to="/admin/users/0">
                        <div class="text-white text-center me-2 d-flex align-items-center justify-content-center">
                        <i class="material-icons opacity-10">person_add</i>
                        </div>
                        <span class="nav-link-text ms-1">ユーザー登録</span>
                    </router-link>
                </li>
                <li class="nav-item">
                    <router-link class="nav-link text-white" to="/admin/users">
                        <div class="text-white text-center me-2 d-flex align-items-center justify-content-center">
                        <i class="material-icons opacity-10">people</i>
                        </div>
                        <span class="nav-link-text ms-1">ユーザー一覧</span>
                    </router-link>
                </li>
            </ul>
        </div>
    </aside>
</template>

<script>
export default {
    name: "DashBoard"
}
</script>