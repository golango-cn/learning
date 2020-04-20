<template>
  <div>

    <el-breadcrumb separator="/">
      <el-breadcrumb-item :to="{ path: '/' }">首页</el-breadcrumb-item>
      <el-breadcrumb-item>商品管理</el-breadcrumb-item>
      <el-breadcrumb-item>商品列表</el-breadcrumb-item>
    </el-breadcrumb>

    <el-row :gutter="20">
      <el-col :span="10">
        <el-input clearable @clear="getGoods" v-model="queryInfo.query" placeholder="请输入内容">
          <el-button slot="append" icon="el-icon-search" @click="searchGoods()"></el-button>
        </el-input>
      </el-col>
      <el-col :span="4">
        <el-button type="primary" @click="addGoods">添加商品</el-button>
      </el-col>
    </el-row>

    <el-table border stripe :data="goods">
      <el-table-column type="index"></el-table-column>
      <el-table-column label="商品名称" prop="name"></el-table-column>
      <el-table-column label="商品价格" width="100px" prop="price"></el-table-column>
      <el-table-column label="商品重量" width="100px" prop="weight"></el-table-column>

      <el-table-column label="创建时间" width="160px" prop="create_at">
        <template slot-scope="scope">
          {{scope.row.create_at|dateFromat}}
        </template>
      </el-table-column>
      <el-table-column width="160px" label="操作">
        <template slot-scope="scope">
          <el-button type="primary" size="mini" @click="editGoods(scope.row.id)">编辑</el-button>
          <el-button type="danger" size="mini" @click="deleteGoods(scope.row.id)">删除</el-button>
        </template>
      </el-table-column>
    </el-table>

    <el-pagination
      @size-change="handleSizeChange"
      @current-change="handleCurrentChange"
      :current-page="queryInfo.pageIndex"
      :page-sizes="[10, 20, 50]"
      :page-size="100"
      layout="total, sizes, prev, pager, next, jumper"
      :total="queryInfo.totalCount">
    </el-pagination>

  </div>
</template>

<script>
  export default {
    data() {
      return {

        queryInfo: {
          pageIndex: 1,
          pageSize: 10,
          query: '',
          totalCount: 0
        },

        // 表单
        form: {},
        // 验证规则
        formRules: {},

        dialogVisible: false,

        goods: []
      }
    },
    created() {
      this.getGoods()
    },
    methods: {
      getGoods: async function() {
        const url = `/goods?pageIndex=${this.queryInfo.pageIndex}&pageSize=${this.queryInfo.pageSize}`
        const { data: resp } = await this.$http.post(url, { keys: this.queryInfo.query })
        console.log(resp)

        this.goods = resp.goods
        this.queryInfo.totalCount = resp.totalCount

      },

      searchGoods: function() {
        this.getGoods()
      },

      handleSizeChange: function(size) {
        this.queryInfo.pageSize = size
        this.getGoods()
      },

      handleCurrentChange: function(index) {
        this.queryInfo.pageIndex = index
        this.getGoods()
      },

      // 编辑商品
      editGoods: function(id) {
        this.$message.warning('方法未实现，商品编号：' + id)
      },

      // 删除商品f
      deleteGoods: function(id) {
        this.$message.warning('方法未实现，商品编号：' + id)
      },

      // 添加商品
      addGoods: function() {
        this.$router.push('goods/edit')
      }
    },

    computed: {
      // 计算列，使用方式
      // <div>{{fileLength}}</div>
      fileLength: function() {
        return this.goods.length
      }
    }
  }
</script>

<style type="less" scoped>

</style>
