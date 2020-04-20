<template>
  <div>
    <el-card>
      <el-form :model="form" :rules="formRules" ref="formRef" label-width="100px">
        <el-form-item label="省市区县" prop="address1">
          <el-cascader :options="cityData"
                       v-model="address1"
                       style="width: 400px"
                       clearable
                       :props="{ expandTrigger: 'hover', checkStrictly: true, label:'name', value: 'code', children:'sub' }"
          ></el-cascader>
        </el-form-item>
      </el-form>

      <el-timeline :reverse="reverse">
        <el-timeline-item
          v-for="(activity, index) in activities"
          :key="index"
          :timestamp="activity.timestamp">
          {{activity.content}}
        </el-timeline-item>
      </el-timeline>


    </el-card>
  </div>
</template>

<script>

  import cityData from './city-data.js'

  export default {
    data() {
      return {
        // 表单
        form: {},
        // 验证规则
        formRules: {
          address1: [{ required: true, message: '请选择省市区县', trigger: 'blur' }]
        },
        address1: [],

        cityData,

        reverse: true,
        activities: [{
          content: '活动按期开始',
          timestamp: '2018-04-15'
        }, {
          content: '通过审核',
          timestamp: '2018-04-13'
        }, {
          content: '创建成功',
          timestamp: '2018-04-11'
        }]
      }
    },
    created() {

    },
    methods: {}
  }
</script>

<style type="less" scoped>

</style>
