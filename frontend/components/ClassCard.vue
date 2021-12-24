<template>
  <div
    class="flex justify-between items-center w-full h-24 mb-5 bg-slate-300 px-10 rounded-lg"
  >
    <p class="text-xl">{{ class_id }}</p>
    <p class="text-xl">{{ class_name }}</p>
    <div class="flex gap-6">
      <button
        class="w-20 h-10 bg-green-500 rounded-lg text-white"
        @click="addClick(class_id, class_name)"
      >
        เพิ่ม
      </button>
      <button
        class="w-20 h-10 bg-blue-600 rounded-lg text-white"
        @click="editClick(class_id)"
      >
        เเก้ไข
      </button>
      <button
        class="w-20 h-10 bg-red-600 rounded-lg text-white"
        @click="deleteClick(class_id)"
      >
        ลบ
      </button>
    </div>
  </div>
</template>

<script>
export default {
  props: ['class_name', 'class_id', 'class_instructor', 'class_detail'],
  methods: {
    addClick(id, className) {
      this.$router.push({
        name: 'admin-lesson',
        query: { classroom: id, classname: className },
      })
    },

    editClick(id) {
      this.$router.push({
        name: 'admin-classroom-editClass',
        query: { classroom: id },
      })
    },

    async deleteClick(id) {
      try {
        if (
          confirm(
            `ต้องการลยวิชา ${this.$props.class_id} : ${this.$props.class_name} ใช่หรือไม่?`
          )
        ) {
          await this.$axios.delete('/classroom/deleteclassroom', {
            params: {
              classroom: id,
            },
          })
          location.reload()
        }
      } catch (error) {}
    },
  },
}
</script>
