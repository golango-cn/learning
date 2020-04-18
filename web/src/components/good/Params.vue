<template>

  <div>
    <el-breadcrumb separator="/">
      <el-breadcrumb-item :to="{ path: '/' }">首页</el-breadcrumb-item>
      <el-breadcrumb-item>商品管理</el-breadcrumb-item>
      <el-breadcrumb-item>分类参数</el-breadcrumb-item>
    </el-breadcrumb>

    <el-card>
      <el-alert
        show-icon
        :closable="false"
        title="警告提示的文案"
        type="warning">
      </el-alert>

      <el-row class="cat_opt">
        <el-col>
          <span>选择商品分类：</span>
          <el-cascader
            v-model="form.cate_pid"
            :options="parentCates"
            clearable
            @change="cateChanged"
            :props="{ expandTrigger: 'hover', label:'cate_name', value: 'id' }">
          </el-cascader>
        </el-col>
      </el-row>

      <el-tabs v-model="activeName" @tab-click="handleClicked">
        <el-tab-pane label="动态参数" name="first">
          <el-button type="primary" size="mini" @click="dialogVisible = true"
                     :disabled="this.form.cate_pid.length !== 3">添加参数
          </el-button>

          <el-table :data="firstData" border>
            <el-table-column type="expand">
              <template slot-scope="scope">
                <el-tag closable :key="item"
                        @close="handleTagClose(scope.row, i)" v-for="(item,i) in scope.row.tags">
                  {{item}}
                </el-tag>

                <el-input
                  class="input-new-tag"
                  v-if="scope.row.inputVisable"
                  v-model="scope.row.inputValue"
                  ref="saveTagInput"
                  size="small"
                  style="width:120px"
                  @keyup.enter.native="handleInputConfirm(scope.row)"
                  @blur="handleInputConfirm(scope.row)"
                >
                </el-input>
                <el-button v-else class="button-new-tag" size="small" @click="showInput(scope.row)">+ New Tag
                </el-button>

              </template>
            </el-table-column>
            <el-table-column label="#" type="index"></el-table-column>
            <el-table-column prop="name" label="名称"></el-table-column>
            <el-table-column label="操作">
              <el-button type="primary" size="mini">编辑</el-button>
              <el-button type="danger" size="mini">删除</el-button>
            </el-table-column>
          </el-table>

        </el-tab-pane>
        <el-tab-pane label="静态属性" name="second">
          <el-button type="primary" size="mini" @click="dialogVisible = true"
                     :disabled="this.form.cate_pid.length !== 3">添加属性
          </el-button>

          <el-table :data="secondData" border>
            <el-table-column type="expand">
              <template slot-scope="scope">
                <el-tag closable :key="item" v-for="(item) in scope.row.tags">
                  {{item}}
                </el-tag>
              </template>
            </el-table-column>
            <el-table-column label="#" type="index"></el-table-column>
            <el-table-column prop="name" label="名称"></el-table-column>
            <el-table-column label="操作">
              <el-button type="primary" size="mini">编辑</el-button>
              <el-button type="danger" size="mini">删除</el-button>
            </el-table-column>
          </el-table>

        </el-tab-pane>
      </el-tabs>

    </el-card>

    <el-dialog
      :title="'添加'+addText()"
      :visible.sync="dialogVisible"
      @close="addFormClosed"
      width="50%">

      <el-form ref="formRef" :model="addForm" :rules="addFormRule" label-width="80px">
        <el-form-item :label="addText()" prop="name">
          <el-input v-model="addForm.name"></el-input>
        </el-form-item>
      </el-form>

      <span slot="footer" class="dialog-footer">
        <el-button @click="dialogVisible = false">取 消</el-button>
        <el-button type="primary" @click="saveParams">确 定</el-button>
      </span>
    </el-dialog>
  </div>
</template>

<script>

  // 验证规则
  export default {
    data() {
      return {
        parentCates: [],
        // 表单
        form: {
          cate_pid: []
        },
        // 验证规则
        formRules: {},
        addForm: {
          name: ''
        },
        activeName: 'first',

        addFormRule: {
          name: [
            { required: true, message: '请输入参数名称', trigger: 'blur' }
          ]
        },

        dialogVisible: false,

        firstData: [],

        secondData: []
      }
    },
    created() {
      this.getParentCates()
    },
    methods: {
      async getParentCates() {
        const { data: resp } = await this.$http.get('/good/cates')
        console.log(resp)
        this.parentCates = resp.cates
      },
      async getAttrs() {
        const { data: resp } = await this.$http.post(`/good/attrs/${this.activeName}`, {
          pid: this.form.cate_pid.join('')
        })

        resp.forEach((item) => {
          item.inputVisable = false
          item.inputValue = ''
        })

        console.log(resp)
        if (this.activeName === 'first') {
          this.firstData = resp
        } else {
          this.secondData = resp
        }
      },
      cateChanged() {
        if (this.form.cate_pid.length !== 3) {
          this.form.cate_pid = []
          return
        }

        // // TODO 通过cate_pid和activeName发起网络请求
        // console.log(this.form.cate_pid, this.activeName, this.cate_pid)
        // this.$message.warning('暂未实现')

        this.getAttrs()

      },
      handleClicked() {

        // // TODO 通过cate_pid和activeName发起网络请求
        // console.log(this.form.cate_pid, this.activeName)
        // this.$message.warning('暂未实现')

        this.getAttrs()

      },
      addText() {
        if (this.activeName === 'first') {
          return '动态参数'
        }
        return '静态属性'
      },
      saveParams() {
        this.$refs.formRef.validate(valid => {
          if (!valid) {
            return
          }

          // TODO 发起网络请求
          console.log('name=', this.addForm.name)
          console.log('activeName=', this.activeName)
          console.log('cate_pid=', this.form.cate_pid, this.form.cate_pid.length)

          this.$message.warning('暂未实现')

          this.$refs.formRef.resetFields()
          this.dialogVisible = false
        })
      },
      addFormClosed() {
        this.$refs.formRef.resetFields()
        console.log('Dialog closed')
      },
      handleInputConfirm(row) {
        if (row.inputValue.trim() === '') {
          row.inputValue = ''
          row.inputVisable = false
          return
        }

        // 有数据输入
        row.tags = row.tags || []
        row.tags.push(row.inputValue.trim())
        row.inputValue = ''
        row.inputVisable = false

        // TODO 发送网络请求保存数据

      },
      showInput(row) {
        row.inputVisable = true
        this.$nextTick(_ => {
          this.$refs.saveTagInput.$refs.input.focus()
        })
      },

      handleTagClose(row, i) {

        row.tags.splice(i, 1)
        // TODO 保存数据库
      }
    }
  }
</script>

<style type="less" scoped>
  .cat_opt {
    margin: 15px 0px;
  }

  .el-tag {
    margin: 5px;
  }

  .input-new-tag {
    margin: 5px;
  }

  .button-new-tag {
    margin: 5px;
  }
</style>
