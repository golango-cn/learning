<template>
  <div>
    <el-breadcrumb separator="/">
      <el-breadcrumb-item :to="{ path: '/home' }">首页</el-breadcrumb-item>
      <el-breadcrumb-item>权限管理</el-breadcrumb-item>
      <el-breadcrumb-item>用户列表</el-breadcrumb-item>
    </el-breadcrumb>

    <el-card>

      <el-table border stripe :data="rights">
        <el-table-column type="index"></el-table-column>
        <el-table-column label="名称" prop="right_name"></el-table-column>
        <el-table-column label="路径" prop="path"></el-table-column>
        <el-table-column label="等级" prop="level"></el-table-column>
        <el-table-column label="操作" width="180px">
          <template slot-scope="scope">
            <el-tag v-if="scope.row.level === '一级'" type="success">一级</el-tag>
            <el-tag v-if="scope.row.level === '二级'" type="warning">二级</el-tag>
            <el-tag v-if="scope.row.level === '三级'" type="error">三级</el-tag>
          </template>
        </el-table-column>
      </el-table>

    </el-card>

  </div>
</template>

<script>
  export default {
    data() {
      return {
        // 表单
        form: {},
        // 验证规则
        formRules: {},
        // 权限列表
        rights: []
      }
    },
    created() {
      this.getRights()
    },
    methods: {
      async getRights() {
        const { data: resp } = await this.$http.post('/rights')
        this.rights = resp.rights
      }
    }
  }
</script>

<style type="less" scoped>

</style>
