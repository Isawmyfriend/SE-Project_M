<template>
  <div>
    <NavigationBar />
    <div class="flex flex-col h-full justify-center pt-16 px-80">
      <div
        class="flex gap-3 h-16 w-full bg-slate-200 rounded-lg mb-14 items-center px-10"
      >
        <a href="/admin">หน้าเเรก</a>
        <p>/</p>
        <a href="/admin/classroom">สร้างห้องเรียน</a>
        <p>/</p>
        <a href="" @click="navBackClick">สร้างบทเรียน</a>
        <p>/</p>
        <p class="text-green-600">เขียนบทเรียน</p>
      </div>

      <div class="flex justify-between items-center mb-10">
        <div class="pb-3">
          <div class="flex items-center gap-3">
            <p class="mb-2">
              บทที่ {{ lessonno }} :
              <span class="text-orange-500 uppercase font-semibold">{{
                lessonname
              }}</span>
            </p>

            <button @click="editClick">
              <svg
                xmlns="http://www.w3.org/2000/svg"
                class="h-8 w-8 pb-2 stroke-gray-600"
                fill="none"
                viewBox="0 0 24 24"
                stroke="currentColor"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z"
                />
              </svg>
            </button>
          </div>

          <p>{{ timeStamp }}</p>
        </div>
        <div class="flex gap-3 items-center">
          <button
            v-show="example"
            class="flex justify-center items-center gap-2 bg-green-600 rounded w-32 h-10"
            @click="exampleClick"
          >
            <svg
              xmlns="http://www.w3.org/2000/svg"
              class="h-6 w-6 stroke-white"
              fill="none"
              viewBox="0 0 24 24"
              stroke="currentColor"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"
              />
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M2.458 12C3.732 7.943 7.523 5 12 5c4.478 0 8.268 2.943 9.542 7-1.274 4.057-5.064 7-9.542 7-4.477 0-8.268-2.943-9.542-7z"
              />
            </svg>
            <p class="text-white" href="">ดูตัวอย่าง</p>
          </button>

          <button
            class="w-20 h-10 ml-4 bg-orange-500 rounded text-white"
            @click="sendFile"
          >
            <span>บันทึก</span>
          </button>
        </div>
      </div>

      <div
        class="h-full w-full shadow-[0_-5px_35px_rgba(0,0,0,0.15)] rounded-xl bottom-0 px-20 py-16"
      >
        <div class="flex flex-col w-full pt-5">
          <div class="flex justify-start">
            <div class="mb-3 w-96">
              <div class="flex gap-3 mb-3">
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  class="h-6 w-6"
                  fill="none"
                  viewBox="0 0 24 24"
                  stroke="orange"
                >
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M4 16v1a3 3 0 003 3h10a3 3 0 003-3v-1m-4-8l-4-4m0 0L8 8m4-4v12"
                  />
                </svg>

                <label
                  for="formFile"
                  class="form-label inline-block mb-2 text-orange-500"
                  >{{ fileStatus }}</label
                >
              </div>

              <input
                class="form-control block w-full px-3 py-1.5 text-base font-normal text-gray-700 bg-white bg-clip-padding border border-solid border-gray-300 rounded transition ease-in-out m-0 focus:text-gray-700 focus:bg-white focus:border-blue-600 focus:outline-none"
                type="file"
                accept=".md"
                @change="handleFileUpload($event)"
              />
            </div>
          </div>

          <div class="flex justify-center mt-9">
            <div class="container w-full h-[48rem] overflow-scroll">
              <article class="prose lg:prose-xl max-w-none pointer-events-none">
                <nuxt-content :document="doc" />
              </article>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  middleware: 'auth',
  data() {
    return {
      classroomname: '',
      lessonname: this.$route.query.lessonname,
      lessonno: this.$route.query.lesson,
      file: '',
      doc: [],
      fileStatus: 'อัพโหลดไฟล์ Markdown',
      filename: '',
      timeStamp: '',
      example: false,
    }
  },
  async created() {
    const res = await this.$axios.get('/blog/getblogbylessonno', {
      params: {
        classroom: this.$route.query.classroom,
        lesson: this.$route.query.lesson,
      },
    })
    this.classroomname = res.data.classroom
    try {
      const file = res.data.blog.lesson_detail
      const doc = await this.$content('blog', file).fetch()
      this.fileStatus = 'ไฟล์ Markdown ล่าสุด : ' + file + '.mb'
      this.filename = file
      this.doc = doc
      this.example = true

      this.timeStamp = 'อัปโหลดล่าสุด : ' + res.data.blog.create_date
    } catch (error) {
      this.doc = []
    }
  },

  methods: {
    handleFileUpload(event) {
      this.file = event.target.files[0]
    },
    async sendFile() {
      try {
        const formData = new FormData()
        formData.append('document', this.file)

        await this.$axios.post('/blog/uploadFile', formData, {
          params: {
            classroom: this.$route.query.classroom,
            lesson: this.$route.query.lesson,
          },
        })

        if (this.filename !== '') {
          await this.$axios.delete('/lesson/deleteFile', {
            params: {
              filename: this.filename,
            },
          })
        }

        alert('เพิ่มไฟล์เรียบร้อยแล้ว')
        location.reload()
      } catch (error) {
        alert('เกิดข้อผิดพลาด')
      }
    },

    exampleClick() {
      this.$router.push({
        name: 'admin-blog-viewBlog',
        query: {
          classroom: this.$route.query.classroom,
          lesson: this.$route.query.lesson,
        },
      })
    },

    editClick() {
      this.$router.push({
        name: 'admin-blog-editBlog',
        query: {
          classroom: this.$route.query.classroom,
          lesson: this.$route.query.lesson,
        },
      })
    },

    navBackClick() {
      this.$router.push({
        name: 'admin-lesson',
        query: {
          classroom: this.$route.query.classroom,
          classname: this.classroomname,
        },
      })
    },
  },
}
</script>
