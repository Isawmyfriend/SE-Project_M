<template>
  <div
    class="fixed inset-0 bg-gray-800 bg-opacity-70 flex items-center justify-center"
  >
    <div class="w-[40rem] h-[48rem] bg-slate-100 p-14 rounded-lg shadow-xl">
      <p class="text-lg font-semibold mb-8">เพิ่มชั้นเรียน</p>
      <div class="mb-6 w-full">
        <div class="flex flex-row justify-between mb-2">
          <p>รหัสวิชา</p>
        </div>
        <input
          v-model="class_id"
          class="border-solid border rounded-md border-slate-300 h-10 w-full px-4 py-2"
          type="text"
        />
      </div>
      <div class="mb-6 w-full">
        <div class="flex flex-row justify-between mb-2">
          <p>ชื่อวิชา</p>
        </div>
        <input
          v-model="class_name"
          class="border-solid border rounded-md border-slate-300 h-10 w-full px-4 py-2"
          type="text"
        />
      </div>
      <div class="mb-6 w-full">
        <div class="flex flex-row justify-between mb-2">
          <p>ผู้สอน</p>
        </div>
        <input
          v-model="class_instructor"
          class="border-solid border rounded-md border-slate-300 h-10 w-full px-4 py-2"
          type="text"
        />
      </div>
      <div class="mb-6 w-full">
        <div class="flex flex-row justify-between mb-2">
          <p>ลิงค์รูปภาพหน้าปก</p>
        </div>
        <input
          v-model="class_image"
          class="border-solid border rounded-md border-slate-300 h-10 w-full px-4 py-2"
          type="text"
        />
      </div>
      <div class="mb-6 w-full">
        <div class="flex flex-row justify-between mb-2">
          <p>คำอธิบายเกี่ยวกับเนื้อหาของวิชา</p>
        </div>

        <textarea
          v-model="class_detail"
          class="border-solid border rounded-md border-slate-300 h-32 w-full px-4 py-2"
        ></textarea>
      </div>
      <div class="flex flex-row-reverse gap-5">
        <button @click="onPopup" class="w-20 h-10 text-black">
          <span>ยกเลิก</span>
        </button>
        <button
          @click="createClick"
          class="w-20 h-10 bg-orange-500 rounded-lg text-white"
        >
          <span>เพิ่ม</span>
        </button>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      class_id: '',
      class_name: '',
      class_instructor: '',
      class_detail: '',
      class_image: '',
    }
  },

  methods: {
    onPopup() {
      this.$emit('notify')
    },

    async createClick() {
      try {
        await this.$axios.post('/classroom/sign-up', {
          class_id: this.class_id,
          class_name: this.class_name,
          class_instructor: this.class_instructor,
          class_detail: this.class_detail,
          class_image: this.class_image,
        })
        alert('เพิ่มชั้นเรียนเรียบร้อยแล้ว')
        location.reload()
      } catch (error) {
        alert('เกิดข้อผิดพลาด')
      }
    },
  },
}
</script>
