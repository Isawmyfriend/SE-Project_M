<template>
  <div
    class="fixed inset-0 bg-slate-300 bg-opacity-70 flex items-center justify-center"
  >
    <div class="w-[40rem] h-[26rem] bg-slate-100 p-14 rounded-lg shadow-xl">
      <p class="text-lg font-semibold mb-8">แก้ไขบทเรียน</p>
      <div class="mb-6 w-full">
        <div class="flex flex-row justify-between mb-2">
          <p>บทที่</p>
        </div>
        <input
          v-model="lesson_no"
          class="border-solid border rounded-md border-slate-300 h-10 w-full px-4 py-2"
          type="number"
          min="1"
        />
      </div>
      <div class="mb-6 w-full">
        <div class="flex flex-row justify-between mb-2">
          <p>ชื่อบท</p>
        </div>
        <input
          v-model="lesson_name"
          class="border-solid border rounded-md border-slate-300 h-10 w-full px-4 py-2"
          type="text"
        />
      </div>

      <div class="flex flex-row-reverse gap-5">
        <button class="w-20 h-10 text-black" @click="onClick('cancel')">
          <span>ยกเลิก</span>
        </button>
        <button
          class="w-20 h-10 bg-orange-500 rounded-lg text-white"
          @click="createClick"
        >
          <span>บันทึก</span>
        </button>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      lesson_no: '',
      lesson_name: '',
      lesson_subject: '',
    }
  },

  async created() {
    try {
      const res = await this.$axios.get('/blog/getblogbylessonno', {
        params: {
          classroom: this.$route.query.classroom,
          lesson: this.$route.query.lesson,
        },
      })

      this.lesson_no = res.data.blog.lesson_no
      this.lesson_name = res.data.blog.lesson_name
      this.lesson_subject = res.data.blog.lesson_subject
    } catch (error) {}
  },

  methods: {
    onClick(action) {
      if (action === 'cancel') {
        window.history.go(-1)
      }
    },

    async createClick() {
      try {
        await this.$axios.post(
          '/blog/updateblog',
          {
            lesson_no: parseInt(this.lesson_no),
            lesson_name: this.lesson_name,
          },
          {
            params: {
              classroom: this.$route.query.classroom,
              lesson: this.$route.query.lesson,
            },
          }
        )
        alert('แก้ไขบทเรียนเรียบร้อยแล้ว')
        this.$router.push('/admin/classroom')
      } catch (error) {
        alert('เกิดข้อผิดพลาด')
      }
    },
  },
}
</script>
