<template>
  <div>
    <div
      class="fixed inset-0 bg-slate-300 bg-opacity-70 flex items-center justify-center"
    >
      <div class="w-[40rem] h-[48rem] bg-slate-100 p-14 rounded-lg shadow-xl">
        <p class="text-lg font-semibold mb-8">แก้ไขชั้นเรียน</p>
        <div class="mb-6 w-full">
          <div class="flex flex-row justify-between mb-2">
            <p>รหัสวิชา</p>
          </div>
          <input
            v-model="classroom.class_id"
            class="border-solid border rounded-md border-slate-300 h-10 w-full px-4 py-2"
            type="text"
          />
        </div>
        <div class="mb-6 w-full">
          <div class="flex flex-row justify-between mb-2">
            <p>ชื่อวิชา</p>
          </div>
          <input
            v-model="classroom.class_name"
            class="border-solid border rounded-md border-slate-300 h-10 w-full px-4 py-2"
            type="text"
          />
        </div>
        <div class="mb-6 w-full">
          <div class="flex flex-row justify-between mb-2">
            <p>ผู้สอน</p>
          </div>
          <input
            v-model="classroom.class_instructor"
            class="border-solid border rounded-md border-slate-300 h-10 w-full px-4 py-2"
            type="text"
          />
        </div>
        <div class="mb-6 w-full">
        <div class="flex flex-row justify-between mb-2">
          <p>ลิงค์รูปภาพหน้าปก</p>
        </div>
        <input
          v-model="classroom.class_image"
          class="border-solid border rounded-md border-slate-300 h-10 w-full px-4 py-2"
          type="text"
        />
      </div>
        <div class="mb-6 w-full">
          <div class="flex flex-row justify-between mb-2">
            <p>คำอธิบายเกี่ยวกับเนื้อหาของวิชา</p>
          </div>

          <textarea
            v-model="classroom.class_detail"
            class="border-solid border rounded-md border-slate-300 h-32 w-full px-4 py-2"
          ></textarea>
        </div>
        <div class="flex flex-row-reverse gap-5">
          <button @click="onClick('cancel')" class="w-20 h-10 text-black">
            <span>ยกเลิก</span>
          </button>
          <button
            @click="editClassroom"
            class="w-20 h-10 bg-orange-500 rounded-lg text-white"
          >
            <span>แก้ไข</span>
          </button>
        </div>
      </div>

      <div class="absolute"></div>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      idBeforeEdit: '',
      classroom: [],
    }
  },
  async created() {
    const res = await this.$axios.get('/classroom/getallclassroom')
    const classroom = res.data.classroom
    const idClassroom = this.$route.query.classroom
    this.idBeforeEdit = idClassroom
    this.classroom = classroom.find((item) => item.class_id === idClassroom)
  },
  methods: {
    async editClassroom() {
      try {
        await this.$axios.post(
          '/classroom/updateclassroom',
          {
            class_id: this.classroom.class_id,
            class_name: this.classroom.class_name,
            class_instructor: this.classroom.class_instructor,
            class_detail: this.classroom.class_detail,
            class_image: this.classroom.class_image,
          },
          {
            params: {
              classroom: this.idBeforeEdit,
            },
          }
        )
        alert('แก้ไขสำเร็จ')
        this.$router.push('/admin/classroom')
      } catch (error) {
        alert('เกิดข้อผิดพลาด')
        location.reload()
      }
    },

    onClick(action) {
      if (action === 'cancel') {
        this.$router.push('/admin/classroom')
      }
    },
  },
}
</script>
