<template>
  <!-- กล่อง Container หลัก -->
  <div class="flex items-center w-80 h-auto flex-col space-y-5">
    <!-- กล่อง Container สีขาว -->
    <div class="bg-white rounded-lg shadow-xl w-80 mx-10">
      <!-- กล่อง Label -->
      <div class="flex justify-center items-center relative">
        <img
          :src="class_image"
          class="flex justify-center items-center bg-local rounded-t-lg h-60 w-80 filter brightness-50"
          alt=""
        />
      </div>
      <!-- ข้อความ -->
      <h3
        class="text-gray-900 dark:text-white mt-4 text-base font-medium tracking-tight pl-5"
      >
        วิชา : {{ class_name }}
      </h3>
      <h2
        class="text-gray-400 dark:text-white text-sm font-medium tracking-tight pl-5 pt-1"
      >
        ผู้สอน : {{ class_instructor }}
      </h2>
      <p
        class="text-gray-500 dark:text-gray-400 mt-2 text-sm pl-5 pt-0 pr-5 pb-5"
      >
        {{ class_detail }}
      </p>
    </div>

    <!-- ปุ่มดูเนื้อหา -->
    <div class="w-50 h-10">
      <button
        class="bg-orange-500 w-80 rounded-lg h-12 text-white shadow-lg hover:bg-orange-600"
        @click="viewClick"
      >
        ดูเนื้อหา
      </button>
    </div>
  </div>
</template>

<script>
export default {
  props: [
    'class_id',
    'class_name',
    'class_instructor',
    'class_detail',
    'class_image',
  ],
  methods: {
    viewClick() {
      this.$router.push({
        name: 'admin-blog-viewBlog',
        query: {
          classroom: this.$props.class_id,
          lesson: '1',
        },
      })
    },
  },
}
</script>
