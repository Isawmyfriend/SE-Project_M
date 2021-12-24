<template>
  <div>
    <NavigationBar />
    <ClassSignupCard v-show="isPopup" @notify="onPopup" />

    <div class="flex flex-col h-screen justify-center pt-16 px-80">
      <div
        class="flex gap-3 h-16 w-full bg-slate-200 rounded-lg mb-14 items-center px-10"
      >
        <a href="/admin">หน้าเเรก</a>
        <p>/</p>
        <p class="text-green-600">สร้างห้องเรียน</p>
      </div>

      <div class="flex justify-between mb-10">
        <div>
          <p>ห้องเรียน : {{ classroomCount }} วิชา</p>
          <p>คณะวิทยาศาสตร์ สาขาวิทยาการคอมพิวเตอร์</p>
        </div>
        <div>
          <button
            @click="onPopup"
            class="group bg-white border-2 border-orange-400 flex gap-3 w-full h-full hover:bg-orange-400 p-3 text-center rounded text-white items-center"
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
              สร้างห้องเรียน
            </span>
          </button>
        </div>
      </div>

      <div
        class=" justify-center h-full w-full shadow-[0_-5px_35px_rgba(0,0,0,0.15)] rounded-xl bottom-0 px-20 py-16"
      >
        <div v-if="classroom.length > 0">
          <ClassCard
            v-for="c in classroom"
            :key="c"
            :class_id="c.class_id"
            :class_name="c.class_name"
            :class_instructor="c.class_instructor"
            :class_detail="c.class_detail"
          />
        </div>
        <p v-else class="mt-6 text-center ">ไม่มีบทเรียนในห้องเรียน</p>
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
      isPopup: false,
      classroomCount: 0,
      classroom: [],
    }
  },

  async created() {
    const res = await this.$axios.get('/classroom/getallclassroom')

    if (res.data.classroom) {
      this.classroom = res.data.classroom
      this.classroomCount = this.classroom.length
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
