<template>
  <div class="login_containter">
    <div class="login_box">
      <el-form ref="formRef" :rules="rules" :model="form" label-width="0px" class="login_form">
        <el-form-item prop="name">
          <el-input v-model="form.name" placeholder="请输入用户名" prefix-icon="el-icon-user"></el-input>
        </el-form-item>
        <el-form-item prop="pwd">
          <el-input v-model="form.pwd" placeholder="请输入密码" type="password" prefix-icon="el-icon-lock"></el-input>
        </el-form-item>
        <el-form-item class="btns">
          <el-button type="primary" @click="login">登录</el-button>
          <el-button type="info" @click="resetLogin">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script>
  export default {
    data() {
      return {
        form: {
          name: '',
          pwd: ''
        },
        rules: {
          name: [
            { required: true, message: '请输入用户名', trigger: 'blur' },
            { min: 3, max: 5, message: '用户名长度在3到5个字符', trigger: 'blur' }
          ],
          pwd: [
            { required: true, message: '请输入密码', trigger: 'blur' },
            { min: 6, max: 10, message: '密码长度在6到10个字符', trigger: 'blur' }
          ]
        }
      }
    },
    methods: {
      resetLogin() {
        this.$refs.formRef.resetFields()
      },
      login() {
        this.$refs.formRef.validate(async valid => {
          if (!valid) {
            return
          }
          const { data: resp } = await this.$http.post('login', this.form)
          console.log(resp)
          if (resp.code !== 0) {
            return this.$message.error(resp.message)
          }
          this.$message.success(resp.message)
          window.sessionStorage.setItem('data', resp)
          await this.$router.push('/home')
        })
      }
    }
  }
</script>

<style lang="less" scoped>
  .login_containter {
    background-color: #2b4b6b;
    height: 100%;
  }

  .login_box {
    width: 450px;
    height: 300px;
    background-color: #ffffff;
    border-radius: 3px;
    position: absolute;
    left: 50%;
    top: 50%;
    transform: translate(-50%, -50%);
  }

  .btns {
    display: flex;
    justify-content: flex-end;
  }

  .login_form {
    position: absolute;
    bottom: 0;
    width: 100%;
    padding: 0 20px;
    box-sizing: border-box;
  }
</style>
