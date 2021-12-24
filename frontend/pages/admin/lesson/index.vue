<template>
  <div>
    <NavigationBar />
    <LessonSignupCard v-show="isPopup" @notify="onPopup" />
    <div class="flex flex-col h-screen justify-center pt-16 px-80">
      <div
        class="flex gap-3 h-16 w-full bg-slate-200 rounded-lg mb-14 items-center px-10"
      >
        <a href="/admin">หน้าเเรก</a>
        <p>/</p>
        <a href="/admin/classroom">สร้างห้องเรียน</a>
        <p>/</p>
        <p class="text-green-600">สร้างบท</p>
      </div>

      <div class="flex justify-between mb-10">
        <div>
          <p>
            ห้องเรียน :
            <span class="text-orange-500 font-semibold">{{
              titleClassroom
            }}</span>
          </p>
          <p>
            จำนวนบทเรียน : <span>{{ lessonCount }}</span>
          </p>
        </div>
        <div>
          <button
            class="group bg-white border-2 border-orange-400 flex gap-3 w-full h-full hover:bg-orange-400 p-3 text-center rounded text-white items-center"
            @click="onPopup"
          >
            <svg
              xmlns="http://www.w3.org/2000/svg"
              class="h-6 w-6 group-hover:stroke-white stroke-orange-400"
              fill="none"
              viewBox="0 0 24 24"
              stroke="currentColor"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M12 9v3m0 0v3m0-3h3m-3 0H9m12 0a9 9 0 11-18 0 9 9 0 0118 0z"
              />
            </svg>
            <span class="group-hover:text-white text-orange-400">
              สร้างบทเรียน
            </span>
          </button>
        </div>
      </div>

      <div
        class="justify-center h-full w-full shadow-[0_-5px_35px_rgba(0,0,0,0.15)] rounded-xl bottom-0 px-20 py-16"
      >
        <div v-if="lesson.length > 0">
          <LessonCard
            v-for="l in lesson"
            :key="l"
            :lesson_no="l.lesson_no"
            :lesson_name="l.lesson_name"
            :lesson_detail="l.lesson_detail"
          />
        </div>

        <p v-else class="mt-6 text-center">ไม่มีบทเรียนในห้องเรียน</p>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'IndexPage',
  middleware: 'auth',
  data() {
    return {
      lesson: [],
      titleClassroom: '',
      lessonCount: 0,
      isPopup: false,
    }
  },
  async created() {
    const res = await this.$axios.get('/lesson/getlessonbyclassid', {
      params: {
        classroom: this.$route.query.classroom,
      },
    })

    this.titleClassroom = this.$route.query.classname

    if (res.data.lesson) {
      this.lesson = res.data.lesson
      this.lessonCount = this.lesson.length
    }
  },

  methods: {
    onPopup() {
      if (this.isPopup) {
        this.isPopup = false
      } else {
        this.isPopup = true
      }
    },
  },
}
</script>
