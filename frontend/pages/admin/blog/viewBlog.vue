<template>
  <div>
    <NavigationBar />
    <div class="flex min-h-screen py-1">
      <div class="bg-slate-100 w-[25rem] pl-4 pt-6 h-screen overflow-y-scroll">
        <div
          class="flex items-center h-20 my-4 pl-4 border-b-2 border-slate-400"
        >
          <p class="text-lg text-slate-600">{{ classroom }}</p>
        </div>
        <SidebarMenu
          v-for="b in allblogs"
          :key="b"
          :lesson_name="b.lesson_name"
          :lesson_no="b.lesson_no"
          :lesson_subject="b.lesson_subject"
          @notify="location.reload()"
        />
      </div>
      <div class="pt-16 w-full px-24 h-screen overflow-y-scroll">
        <div v-if="!content" class="prose lg:prose-xl max-w-none pointer-events-none">
          <p class=" text-2xl">Oops.... ตอนนี้ยังไม่มีเนื้อหา 🌵</p>
        </div>
        <article v-else class="prose lg:prose-xl max-w-none pointer-events-none">
          <nuxt-content :document="doc" />
        </article>
        
      </div>
    </div>
  </div>
</template>

<script>
export default {
  middleware: 'auth',
  data() {
    return {
      classroom: '',
      content: false,
      doc: [],
      allblogs: [],
    }
  },
  async created() {
    try {
      const res = await this.$axios.get('/blog/getblogbyclasseoom', {
        params: {
          classroom: this.$route.query.classroom,
        },
      })

      const lessonno = parseInt(this.$route.query.lesson)
      this.classroom = res.data.classroom
      this.allblogs = res.data.blog

      const file = this.allblogs.find(function (item) {
        return item.lesson_no === lessonno
      }).lesson_detail

      const doc = await this.$content('blog', file).fetch()
      this.doc = doc
      this.content = true
    } catch (error) {
  
      this.doc = []
    }
  },
}
</script>
