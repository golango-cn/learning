<template>
  <div>
    <el-breadcrumb separator="/">
      <el-breadcrumb-item :to="{ path: '/home' }">首页</el-breadcrumb-item>
      <el-breadcrumb-item><a href="/users">用户管理</a></el-breadcrumb-item>
      <el-breadcrumb-item>用户列表</el-breadcrumb-item>
    </el-breadcrumb>

    <el-card>

      <el-row :gutter="20">
        <el-col :span="10">
          <el-input clearable @clear="getUsers" v-model="queryInfo.query" placeholder="请输入内容">
            <el-button slot="append" icon="el-icon-search" @click="searchUsers()"></el-button>
          </el-input>
        </el-col>
        <el-col :span="4">
          <el-button type="primary" @click="dialogVisible=true">添加用户</el-button>
        </el-col>
      </el-row>

      <el-table border stripe :data="users">
        <el-table-column type="index"></el-table-column>
        <el-table-column label="姓名" prop="user_name"></el-table-column>
        <el-table-column label="角色" prop="role_name"></el-table-column>
        <el-table-column label="电话" prop="mobile"></el-table-column>
        <el-table-column label="创建日期" prop="create_time"></el-table-column>
        <el-table-column label="状态">
          <template slot-scope="scope">
            <el-switch v-model="scope.row.status" @change="statusChanged(scope.row)"></el-switch>
          </template>
        </el-table-column>
        <el-table-column label="操作" width="180px">
          <template slot-scope="scope">
            <el-button type="primary" size="mini" icon="el-icon-edit" @click="showEditDialog(scope.row.id)"></el-button>
            <el-button type="danger" size="mini" icon="el-icon-delete" @click="deleteUser(scope.row.id)"></el-button>
            <el-tooltip class="item" effect="dark" :enterable="false" content="分配角色" placement="top">
              <el-button type="warning" size="mini" icon="el-icon-setting" @click="setRole(scope.row)"></el-button>
            </el-tooltip>
          </template>
        </el-table-column>
      </el-table>

      <el-pagination
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"
        :current-page="queryInfo.pageIndex"
        :page-sizes="[10, 20, 50, 100]"
        :page-size="queryInfo.pageSize"
        layout="total, sizes, prev, pager, next, jumper"
        :total="totalCount">
      </el-pagination>

    </el-card>

    <el-dialog
      title="添加用户"
      :visible.sync="dialogVisible"
      @close="dialogClosed"
      width="50%">

      <el-form ref="formRef" :model="form" :rules="formRules" label-width="70px">
        <el-form-item label="用户名" prop="name">
          <el-input v-model="form.name"></el-input>
        </el-form-item>
        <el-form-item label="密码" prop="password">
          <el-input v-model="form.password"></el-input>
        </el-form-item>
        <el-form-item label="邮箱" prop="email">
          <el-input v-model="form.email"></el-input>
        </el-form-item>
        <el-form-item label="手机" prop="mobile">
          <el-input v-model="form.mobile"></el-input>
        </el-form-item>
      </el-form>

      <span slot="footer" class="dialog-footer">
        <el-button @click="dialogVisible = false">取 消</el-button>
        <el-button type="primary" @click="submit">确 定</el-button>
      </span>
    </el-dialog>

    <el-dialog
      title="修改用户"
      :visible.sync="editDialogVisible"
      width="50%">

      <el-form ref="editFormRef" :model="form" :rules="formRules" label-width="70px">
        <el-form-item label="用户名" prop="name">
          <el-input v-model="form.name" disabled></el-input>
        </el-form-item>
        <el-form-item label="邮箱" prop="email">
          <el-input v-model="form.email"></el-input>
        </el-form-item>
        <el-form-item label="手机" prop="mobile">
          <el-input v-model="form.mobile"></el-input>
        </el-form-item>
      </el-form>

      <span slot="footer" class="dialog-footer">
        <el-button @click="editDialogVisible = false">取 消</el-button>
        <el-button type="primary" @click="modifySubmit()">确 定</el-button>
      </span>
    </el-dialog>

    <el-dialog
      title="分配角色"
      :visible.sync="dialogRoleVisible"
      width="50%">
      <p>用户名称：{{user.user_name}}</p>
      <p>角色名称：{{user.role_name}}</p>
      <p>
        <el-select v-model="selectedRoleId" placeholder="请选择">
          <el-option
            v-for="item in roles"
            :key="item.value"
            :label="item.label"
            :value="item.value">
          </el-option>
        </el-select>
      </p>
      <span slot="footer" class="dialog-footer">
        <el-button @click="dialogRoleVisible = false">取 消</el-button>
        <el-button type="primary" @click="saveUserRole">确 定</el-button>
      </span>
    </el-dialog>

  </div>
</template>

<script>
  export default {
    data() {

      // 验证邮箱
      var checkEmail = (rule, value, callback) => {
        const regEmail = /^[a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+$/
        if (regEmail.test(value)) {
          return callback()
        }
        return callback(new Error('请输入合法的邮箱'))
      }

      // 验证手机
      var checkMobile = (rule, value, callback) => {
        const regMobile = /^1\d{10}$/
        if (regMobile.test(value)) {
          return callback()
        }
        return callback(new Error('请输入合法的手机号'))
      }

      return {
        // 获取用户列表参数
        queryInfo: {

          // 查询条件
          query: '',

          // 当前页
          pageIndex: 1,

          // 每页条数
          pageSize: 20
        },

        // 用户列表
        users: [],

        // 总数量
        totalCount: 0,

        // 控制添加用户对话框显示与隐藏
        dialogVisible: false,
        // 控制修改对话框显示与隐藏
        editDialogVisible: false,

        // 添加用户表单
        form: {
          name: '',
          password: '',
          email: '',
          mobile: ''
        },

        // 分配角色
        dialogRoleVisible: false,

        // 当前操作用户
        user: {},

        // 所选角色
        selectedRoleId: '',

        // 角色
        roles: [
          { label: '超级管理员', value: 1 },
          { label: '管理员', value: 2 },
          { label: '普通用户', value: 3 }],

        // 验证规则
        formRules: {
          name: [
            { required: true, message: '请输入用户名', trigger: 'blur' },
            { min: 3, max: 10, message: '用户名的长度在3-10个之间', trigger: 'blur' }
          ],
          password: [
            { required: true, message: '请输入密码', trigger: 'blur' },
            { min: 6, max: 10, message: '密码的长度在6-10个之间', trigger: 'blur' }
          ],
          email: [
            { required: true, message: '请输入邮箱', trigger: 'blur' },
            { validator: checkEmail, trigger: 'blur' }
          ],
          mobile: [
            { required: true, message: '请输入手机', trigger: 'blur' },
            { validator: checkMobile, trigger: 'blur' }
          ]
        }

      }
    },
    created() {

      // 获取用户列表
      this.getUsers()
    },
    methods: {

      // 获取用户列表
      async getUsers() {
        const { data: resp } = await this.$http.get('/users', { params: this.queryInfo })
        if (resp.code !== 0) {
          return this.$message.error(resp.message)
        }
        this.users = resp.users
        this.totalCount = resp.totalCount
      },

      // 每页条数
      handleSizeChange(newSize) {
        this.queryInfo.pageSize = newSize
        this.getUsers()
      },

      // 当前页
      handleCurrentChange(newPage) {
        this.queryInfo.pageIndex = newPage
        this.getUsers()
      },

      // 用户状态改变
      statusChanged(userInfo) {
        // TODO修改用户状态
        this.$message.warning(`调用API接口修改用户状态，用户：${userInfo.id} 状态：${userInfo.status}`)
      },

      // 搜索用户
      searchUsers() {
        this.getUsers()
      },

      // 关闭对话框
      dialogClosed() {
        console.log('触发关闭对话框')
        this.$refs.formRef.resetFields()
      },

      // 提交保存
      submit() {
        this.$refs.formRef.validate(async valid => {
          if (!valid) {
            return
          }
          // TODO 校验通过，发起网络请求，新增用户
          // const { data: resp } = await this.$http.post('/users', this.form)
          // console.log(resp)

          this.dialogVisible = false
          this.getUsers()

        })
      },
      // 编辑用户
      async showEditDialog(id) {
        this.editDialogVisible = true
        console.log(`当前用户编号：${id}`)

        // TODO 这里发送API请求，获取用户信息
        // const { data: resp } = await this.$http.get(`/users/${id}`)
        // console.log(resp)

        this.form.email = 'username@email.com'
        this.form.mobile = '15210001111'
        this.form.name = 'username'

      },

      // 修改用户
      async modifySubmit() {

        this.$refs.editFormRef.validate(async valid => {
          if (!valid) {
            return
          }
          // TODO 校验通过，发起网络请求，修改用户
          // const { data: resp } = await this.$http.post('/users', this.form)
          // console.log(resp)

          this.editDialogVisible = false
          this.getUsers()

        })
      },

      // 删除用户
      async deleteUser(id) {

        var confirm = await this.$confirm('此操作将永久删除该用户, 是否继续?', '提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning'
        }).catch(err => err)

        // 确认删除
        if (confirm === 'confirm') {
          // TODO 发起网络请求，删除用户
          this.$message.warning(`删除用户：${id}`)
          this.getUsers()

        } else {
          this.$message('取消删除')
        }

      },

      // 分配角色
      async setRole(user) {
        console.log(user)
        this.user = user
        this.dialogRoleVisible = true
      },

      // 保存角色
      async saveUserRole() {

        console.log('当前用户', this.user.id)
        console.log('所选角色', this.selectedRoleId)

        this.$message.success('操作成功')
        this.dialogRoleVisible = false
      }
    }
  }
</script>

<style type="less" scoped>

</style>
