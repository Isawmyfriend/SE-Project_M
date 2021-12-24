<template>
  <div
    class="flex justify-between items-center w-full h-24 bg-slate-300 px-10 mb-5 rounded-lg"
  >
    <p class="text-xl">{{ lesson_no }}</p>
    <p class="text-xl">{{ lesson_name }}</p>
    <div class="flex gap-6">
      <button
        v-show="!isEditing"
        class="w-20 h-10 bg-green-500 rounded-lg text-white"
        @click="addClick(lesson_no)"
      >
        เพิ่ม
      </button>
      <button
        v-show="isEditing"
        class="w-20 h-10 bg-blue-600 rounded-lg text-white"
        @click="addClick(lesson_no)"
      >
        เเก้ไข
      </button>
      <button
        class="w-20 h-10 bg-red-600 rounded-lg text-white"
        @click="deleteClick(lesson_no)"
      >
        ลบ
      </button>
    </div>
  </div>
</template>

<script>
export default {
  props: ['lesson_no', 'lesson_name', 'lesson_detail'],
  data() {
    return {
      isEditing: false,
      lessonDetail: this.$props.lesson_detail,
    }
  },

  created() {
    if (this.lessonDetail !== '') {
      this.isEditing = true
    }
  },

  methods: {
    addClick(lessonNo) {
      this.$router.push({
        name: 'admin-blog-createBlog',
        query: {
          classroom: this.$route.query.classroom,
          lesson: lessonNo,
          lessonname: this.$props.lesson_name,
        },
      })
    },
    async deleteClick(id) {
      try {
        if (
          confirm(
            `ต้องการลบ บท ${this.$props.lesson_no} : ${this.$props.lesson_name} ใช่หรือไม่?`
          )
        ) {
          await this.$axios.delete('/lesson/deletelesson', {
            params: {
              classroom: this.$route.query.classroom,
              lesson: id,
            },
          })
          location.reload()
        }
      } catch (error) {}
    },
  },
}
</script>
