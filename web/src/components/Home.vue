<template>
  <el-container class="home-container">
    <el-header>
      <div>
        <span>后台管理</span>
      </div>
      <el-button type="info" @click="loginOut">退出</el-button>
    </el-header>
    <el-container>
      <el-aside :width="isCollapse ? '64px':'200px' ">

        <div class="toggle-button" @click="toggleCollapse">|||</div>

        <el-menu
          router
          :default-active="activePath"
          unique-opened
          :collapse="isCollapse"
          :collapse-transition="false"
          background-color="#333744"
          text-color="#fff"
          active-text-color="#409eff">
          <el-submenu :index="item.id+''" v-for="item in menus" :key="item.id">

            <template slot="title">
              <i :class="icons[item.id]"></i>
              <span>{{item.name}}</span>
            </template>

            <el-menu-item :index="subItem.path+''" v-for="subItem in item.children" :key="subItem.id"
                          @click="navState(subItem.path)">
              <i class="el-icon-menu"></i>
              <span>{{subItem.name}}</span>
            </el-menu-item>

          </el-submenu>
        </el-menu>
      </el-aside>
      <el-main>
        <router-view></router-view>
      </el-main>
    </el-container>
  </el-container>

</template>

<script>
  export default {
    data() {
      return {
        menus: [],
        icons: {
          1: 'el-icon-user',
          2: 'el-icon-lock',
          3: 'el-icon-basketball',
          4: 'el-icon-cherry',
          5: 'el-icon-coffee'
        },
        isCollapse: false,
        // 初激活的链接地址
        activePath: ''
      }
    },
    created() {
      this.getMenus()
      this.activePath = window.sessionStorage.getItem('activePath')
    },
    methods: {
      navState(activePath) {
        window.sessionStorage.setItem('activePath', activePath)
        this.activePath = activePath
      },
      loginOut() {
        window.sessionStorage.clear()
        this.$router.push('/login')
      },
      // 获取菜单
      async getMenus() {
        const { data: resp } = await this.$http.get('menus')
        this.menus = resp.menus
      },
      toggleCollapse() {
        this.isCollapse = !this.isCollapse
      }
    }
  }
</script>

<style type="less" scoped>
  .home-container {
    height: 100%;
  }

  .el-header {
    background-color: #373d41;
    display: flex;
    justify-content: space-between;
    padding-left: 10px;
    align-items: center;
    color: #fff;
    font-size: 20px;
  }

  .el-aside {
    background-color: #333744;
  }

  .el-menu {
    border-right: none;
  }

  .el-main {
    background-color: #eaedf1;
  }

  .toggle-button {
    background-color: #4a5064;
    font-size: 10px;
    line-height: 24px;
    color: #fff;
    text-align: center;
    letter-spacing: 0.2em;
    cursor: pointer;
  }
</style>
