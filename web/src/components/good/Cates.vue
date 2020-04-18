<template>
  <div>
    <el-breadcrumb separator="/">
      <el-breadcrumb-item :to="{ path: '/' }">首页</el-breadcrumb-item>
      <el-breadcrumb-item>商品管理</el-breadcrumb-item>
      <el-breadcrumb-item>商品分类</el-breadcrumb-item>
    </el-breadcrumb>

    <el-card>
      <el-row>
        <el-col>
          <el-button type="primary" @click="showAddCateDialog">添加分类</el-button>
        </el-col>
      </el-row>

      <tree-table :data="cates"
                  class="tree-table"
                  :selection-type="false"
                  show-index index-text="#"
                  border
                  :show-row-hover="false"
                  children-prop="children"
                  :expand-type="false"
                  :columns="columns">
        <template slot="isok" slot-scope="scope">
          <i class="el-icon-success" style="color: lightgreen;" v-if="scope.row.cate_delete === false"></i>
          <i class="el-icon-error" style="color: red;" v-else></i>
        </template>

        <template slot="orderBy" slot-scope="scope">
          <el-tag size="mini" v-if="scope.row.cate_level === 0">一级</el-tag>
          <el-tag size="mini" v-else-if="scope.row.cate_level === 1" type="success">二级</el-tag>
          <el-tag size="mini" v-else type="warning">三级</el-tag>
        </template>

        <template slot="operation" slot-scope="scope">
          <el-button type="primary" size="mini">编辑</el-button>
          <el-button type="danger" size="mini">删除</el-button>
        </template>

      </tree-table>
    </el-card>

    <el-dialog
      title="添加分类"
      :visible.sync="dialogVisible"
      @close="cateCloed"
      width="50%">

      <el-form ref="formRef" :model="form" :rules="formRules" label-width="80px">
        <el-form-item label="分类名称" prop="cate_name">
          <el-input v-model="form.cate_name"></el-input>
        </el-form-item>
        <el-form-item label="父级分类" prop="cate_pid">
          <el-cascader
            v-model="form.cate_pid"
            :options="parentCates"
            clearable
            :props="{ expandTrigger: 'hover', checkStrictly: true, label:'cate_name', value: 'id' }"></el-cascader>
        </el-form-item>
      </el-form>

      <span slot="footer" class="dialog-footer">
        <el-button @click="dialogVisible = false">取 消</el-button>
        <el-button type="primary" @click="saveCate">确 定</el-button>
      </span>
    </el-dialog>

  </div>
</template>

<script>
  export default {
    data() {
      return {
        // 表单
        form: {
          cate_name: '',
          cate_pid: [1, 101, 1011],
          cate_level: 0
        },
        // 验证规则
        formRules: {
          cate_name: [
            { required: true, message: '请输入分类名称', trigger: 'blur' }
          ],
          cate_pid: [
            { required: true, message: '请选择父分类', trigger: 'blur' }
          ]
        },

        cates: [],

        parentCates: [],

        columns: [
          {
            label: '分类名称',
            prop: 'cate_name'
          },
          {
            label: '是否有效',
            type: 'template',
            template: 'isok'
          },
          {
            label: '排序',
            type: 'template',
            template: 'orderBy'
          },
          {
            label: '操作',
            type: 'template',
            template: 'operation'
          }
        ],

        dialogVisible: false
      }
    },
    created() {
      this.getCates()
    },
    methods: {
      async getCates() {
        const { data: resp } = await this.$http.get('/good/cates')
        console.log(resp)
        this.cates = resp.cates
      },
      showAddCateDialog() {
        this.getParentCates()
        this.dialogVisible = true
      },

      async saveCate() {

        console.log(123)

        this.$refs.formRef.validate(async valid => {

          if (!valid) {
            return
          }
          //
          // // TODO 发起网络请求保存分类

          console.log(this.form.cate_pid)

          this.$message.success('保存成功')
          await this.getCates()
          this.form = {}

          this.dialogVisible = false

        })

      },
      async getParentCates() {
        const { data: resp } = await this.$http.get('/good/cates')
        console.log(resp)
        this.parentCates = resp.cates
      },
      cateCloed() {
        console.log('Form Closed')
        this.$refs.formRef.resetFields()
      }
    }
  }
</script>

<style type="less" scoped>
  .tree-table {
    margin-top: 10px;
  }

  .el-cascader {
    width: 100%;
  }
</style>
