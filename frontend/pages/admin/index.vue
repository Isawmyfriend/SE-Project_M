<template>
  <div class="bg-gray-50">
    <NavigationBar />
    <div class="px-80 pt-20">
      <div class="flex w-full relative h-96 justify-start items-center">
        <!-- ภาพปก -->

        <div
          class="flex flex-col items-start px-40 py-12 w-full h-[22rem] bg-cover bg-[url('@/assets/image/coverimage1.png')] bg-opacity-50 rounded-lg"
        >
          <p class="text-5xl text-orange-500 font-bold mb-5 mt-2">
            CS <span class="text-white">Concept</span>
          </p>
          <p class="text-2xl text-white font-bold">"Hi! Students"</p>
          <p class="text-base text-white mt-4 mb-9">
            เว็บบล็อกแนะนำบทเรียนของสาขาวิทยาการคอมพิวเตอร์เรา
            <br />
            เราได้ทำการรวบรวมบทเรียนทั้งหมดที่สามารถเข้าถึงได้ทุกที่และตลอดเวลา
          </p>
          <a href="https://www.reg.kmitl.ac.th/index/">
            <button
              class="bg-orange-500 w-48 rounded h-10 text-white shadow-lg hover:bg-orange-600 font-bold text-xl flex flex-row justify-center items-center"
            >
              <span class="text-lg">KMITL.REG</span>

              <span class="m-3"
                ><img src="@/assets/image/kmitl.regLogo.svg" alt=""
              /></span>
            </button>
          </a>
        </div>
      </div>
    </div>
    <div class="flex flex-col justify-center items-center pt-8 pb-20">
      <img src="@/assets/image/Line.svg" class="pt-10" alt="" />
      <span class="text-4xl text-gray-600 font-bold pb-16 pt-10">
        Computer - Science</span
      >
      <div
        class="flex flex-row flex-wrap px-80 justify-center gap-36 items-center"
      >
        <SubjectCard
          v-for="c in classroom"
          :key="c"
          :class_id="c.class_id"
          :class_name="c.class_name"
          :class_instructor="c.class_instructor"
          :class_detail="c.class_detail"
          :class_image="c.class_image"
        />

        <div v-if="!havaClass">
          <p class="text-2xl">ตอนนี้ยังไม่มีห้องเรียน 🥲</p>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import NavigationBar from '~/components/NavigationBar.vue'
import SubjectCard from '~/components/SubjectCard.vue'

export default {
  name: 'IndexPage',
  components: { NavigationBar, SubjectCard },
  middleware: 'auth',
  data() {
    return {
      classroomCount: 0,
      classroom: [],
      havaClass: false,
    }
  },

  async created() {
    const res = await this.$axios.get('/classroom/getallclassroom')

    if (res.data.classroom) {
      this.classroom = res.data.classroom
      this.classroomCount = this.classroom.length
      this.havaClass = true
    }
  },
}
</script>
