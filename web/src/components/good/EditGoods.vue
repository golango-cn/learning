<template>
  <div>
    <el-breadcrumb separator="/">
      <el-breadcrumb-item :to="{ path: '/' }">首页</el-breadcrumb-item>
      <el-breadcrumb-item>商品管理</el-breadcrumb-item>
      <el-breadcrumb-item>添加商品</el-breadcrumb-item>
    </el-breadcrumb>

    <el-card>
      <el-alert title="消息提示的文案" type="info" show-icon :closable="false"></el-alert>

      <el-steps :active="acvityIndex-0" finish-status="success" align-center>
        <el-step title="基本信息"></el-step>
        <el-step title="商品参数"></el-step>
        <el-step title="商品属性"></el-step>
        <el-step title="商品图片"></el-step>
        <el-step title="商品内容"></el-step>
        <el-step title="完成"></el-step>
      </el-steps>

      <el-form :model="addForm" :rules="addFormRule" label-position="top" ref="addFormRef" label-width="100px">
        <el-tabs v-model="acvityIndex" tab-position="left" :before-leave="toggleTab">
          <el-tab-pane label="基本信息" name="0">
            <el-form-item label="商品名称" prop="name">
              <el-input v-model="addForm.name"></el-input>
            </el-form-item>
            <el-form-item label="商品价格" prop="price">
              <el-input v-model="addForm.price" type="number"></el-input>
            </el-form-item>
            <el-form-item label="商品重量" prop="weight">
              <el-input v-model="addForm.weight" type="number"></el-input>
            </el-form-item>
            <el-form-item label="商品数量" prop="number">
              <el-input v-model="addForm.number" type="number"></el-input>
            </el-form-item>

          </el-tab-pane>
          <el-tab-pane label="商品参数" name="1">商品参数</el-tab-pane>
          <el-tab-pane label="商品属性" name="2">商品属性</el-tab-pane>
          <el-tab-pane label="商品图片" name="3">
            <el-upload
              action="/upload"
              :http-request="uploadImg"
              :on-preview="handlePreview"
              :on-remove="handleRemove"
              list-type="picture">
              <el-button size="small" type="primary">点击上传</el-button>
              <div slot="tip" class="el-upload__tip">只能上传jpg/png文件，且不超过500kb</div>
            </el-upload>
          </el-tab-pane>
          <el-tab-pane label="商品内容" name="4">商品内容</el-tab-pane>
        </el-tabs>
      </el-form>

    </el-card>

  </div>
</template>

<script>
  export default {
    data() {
      return {
        // 表单
        addForm: {
          name: '',
          weight: 0,
          number: 0,
          price: 0
        },
        // 验证规则
        addFormRule: {
          name: [{ required: true, message: '请输入商品名称', trigger: 'blur' }],
          price: [{ required: true, message: '请输入商品价格', trigger: 'blur' }],
          number: [{ required: true, message: '请输入商品数量', trigger: 'blur' }],
          weight: [{ required: true, message: '请输入商品重量', trigger: 'blur' }]
        },

        acvityIndex: '0'
      }
    },
    created() {
    },
    methods: {
      toggleTab(activityName, oldActivityName) {
        if (oldActivityName === '0') {
          if (this.addForm.name === '') {
            this.$message.warning('请输入商品名称')
            return false
          }
        }
      },
      handlePreview() {

      },
      handleRemove() {

      },
      async uploadImg(f) {
        const param = new FormData()
        param.append('file', f.file)

        const { data: resp } = this.$http.post('/upload', param, { headers: { 'Content-type': 'multipart/form-data' } })
        console.log(resp)

        console.log(f)
      }
    }
  }
</script>

<style type="less" scoped>

</style>
