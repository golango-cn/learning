<template>
  <div>

    <el-breadcrumb separator="/">
      <el-breadcrumb-item :to="{ path: '/home' }">首页</el-breadcrumb-item>
      <el-breadcrumb-item>角色管理</el-breadcrumb-item>
      <el-breadcrumb-item>角色列表</el-breadcrumb-item>
    </el-breadcrumb>

    <el-card>

      <el-table border stripe :data="roles">
        <el-table-column type="expand">
          <template slot-scope="scope">
            <el-row :class="['vcenter', 'bdbottom', i1 === 0 ? 'bdtop':'']" v-for="(item1, i1) in scope.row.children"
                    :key="item1.id">
              <!-- 一级权限 -->
              <el-col :span="5">
                <el-tag>{{item1.auth_name}}</el-tag>
                <i class="el-icon-caret-right"></i>
              </el-col>
              <el-col :span="19">
                <el-row :class="[i2 === 0 ? '':'bdtop', 'vcenter']" v-for="(item2, i2) in item1.children"
                        :key="item2.id">
                  <!-- 渲染二级权限 -->
                  <el-col :span="6">
                    <el-tag type="success">{{item2.auth_name}}</el-tag>
                    <i class="el-icon-caret-right"></i>
                  </el-col>
                  <el-col :span="18">
                    <el-tag closable @close="removeRightById(item3.id)" type="warning"
                            v-for="(item3) in item2.children"
                            :key="item3.id">
                      {{item3.auth_name}}
                    </el-tag>
                  </el-col>
                </el-row>
              </el-col>
            </el-row>
          </template>
        </el-table-column>
        <el-table-column type="index"></el-table-column>
        <el-table-column label="名称" prop="role_name"></el-table-column>
        <el-table-column label="备注" prop="role_desc"></el-table-column>
        <el-table-column label="操作" width="300px">
          <template slot-scope="scope">
            <el-button type="primary" size="mini" icon="el-icon-edit">编辑</el-button>
            <el-button type="danger" size="mini" icon="el-icon-delete">删除</el-button>
            <el-button type="warning" size="mini" icon="el-icon-setting" @click="setRoleRules(scope.row)">分配权限
            </el-button>
          </template>
        </el-table-column>
      </el-table>

    </el-card>

    <el-dialog
      title="提示"
      :visible.sync="ruleDialogVisible"
      width="50%">

      <el-tree ref="ruleTree" :data="rules" default-expand-all :default-checked-keys="defaultCheckedKeys" node-key="id"
               show-checkbox :props="defaultRuleProps"></el-tree>

      <span slot="footer" class="dialog-footer">
       <el-button @click="ruleDialogVisible = false">取 消</el-button>
       <el-button type="primary" @click="saveRoleRuls()">确 定</el-button>
     </span>
    </el-dialog>

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

        roles: [],

        ruleDialogVisible: false,
        // 树型权限
        rules: [],

        defaultRuleProps: {
          label: 'auth_name',
          children: 'children'
        },
        roleId: '',
        defaultCheckedKeys: [10111, 10112]
      }
    },
    created() {
      this.getRoles()
    },
    methods: {

      // 获取角色列表
      async getRoles() {
        const { data: resp } = await this.$http.post('/roles')
        this.roles = resp.roles
      },

      // 根据id删除权限
      async removeRightById(id) {
        const confirm = await this.$confirm('此操作将永久删除该权限, 是否继续?', '提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning'
        }).catch(err => err)

        if (confirm === 'confirm') {
          // TODO 发起网络请求，删除
          this.$message.success(`删除成功：${id}`)
          this.getRoles()
        } else {
          // 取消删除
          this.$message.info('取消了删除')
        }
      },

      // 设置角色权限
      async setRoleRules(role) {
        this.roleId = role.id
        const { data: resp } = await this.$http.get('/rules', { params: role.id })
        this.rules = resp
        console.log(resp)
        this.ruleDialogVisible = true
      },

      // 保存角色权限
      async saveRoleRuls() {
        const keys = [...this.$refs.ruleTree.getCheckedKeys(), ...this.$refs.ruleTree.getHalfCheckedKeys()]
        console.log(keys)
        console.log(this.roleId)

        // TODO 发送http请求，保存角色权限

        this.$message.success('操作成功')
        this.getRoles()
        this.ruleDialogVisible = false

      }
    }
  }
</script>

<style type="less" scoped>
  .el-tag {
    margin: 7px;
  }

  .bdtop {
    border-top: 1px solid #eee;
  }

  .bdbottom {
    border-bottom: 1px solid #eee;
  }

  .vcenter {
    display: flex;
    align-items: center;
  }
</style>
