<template>
  <div>
    <div
      v-if="isDisplay"
      @click="navClick"
      class="flex items-center my-3 h-[4.5rem] pl-4 py-5 bg-white border-t-2 border-l-2 border-b-2 border-gray-400 pointer-events-none"
    >
      <p class="text-lg text-slate-600">{{ lesson_name }}</p>
    </div>
    <div
      v-else
      @click="navClick"
      class="flex items-center my-3 h-[4.5rem] pl-4 py-5 hover:bg-slate-200 cursor-pointer"
    >
      <p class="text-lg text-slate-600">{{ lesson_name }}</p>
    </div>
  </div>
</template>

<script>
export default {
  props: ['lesson_name', 'lesson_no', 'lesson_subject'],
  data() {
    return {
      isDisplay: false,
    }
  },
  created() {
    if (this.$props.lesson_no === parseInt(this.$route.query.lesson)) {
      this.isDisplay = true
    }
  },
  methods: {
    navClick() {
      console.log(this.$props.lesson_subject)
      this.$router.push({
        name: 'admin-blog-viewBlog',
        query: {
          classroom: this.$props.lesson_subject,
          lesson: this.$props.lesson_no,
        },
      })
      this.$emit('notify')
    },
  },
}
</script>
