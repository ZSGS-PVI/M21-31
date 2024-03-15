<!-- <template>
  <div class="container">
    <div class="row">
      <div class="col-md-12">
        <!-- Original Table with "MORE" button -->
        <el-table :data="visibleData" style="width: 100%">
          <el-table-column v-for="(field, index) in visibleFields.slice(0, 5)" :key="index" :prop="field.field" :label="field.label"></el-table-column>
          <el-table-column label="MORE">
            <template slot-scope="scope">
              <el-button @click="showIndividualData(scope.row)" type="text" icon="el-icon-more"></el-button>
            </template>
          </el-table-column>
        </el-table>
        
        <el-drawer :visible="drawerVisible" direction="rtl" size="40%" @close="drawerVisible = false">
          <div style="padding: 20px;">
            <el-table :data="drawerData" style="width: 100%">
              <el-table-column prop="key" label="Element"></el-table-column>
              <el-table-column prop="value" label="Value"></el-table-column>
            </el-table>
          </div>
        </el-drawer>

        <div v-if="activeTab === 1" class="button-container">
          <button @click="prevPage" class="btn btn-primary button">Prev</button>
          <button @click="nextPage" class="btn btn-primary button">Next</button>
        </div>

      </div>
    </div>
  </div>
</template>

<script>
import TabItem from '../components/TabItem.vue'

export default {
  components: {
    TabItem
  },
  props: ['columns', 'liveTableData', 'logsTableData'],
  data() {
    return {
      activeTab: 0,
      offset: 0,
      drawerVisible: false,
      drawerData: []
    };
  },
  computed: {
    visibleFields() {
      return this.columns.map(field => ({ ...field }));
    }
  },
  watch: {
    activeTab() {
      this.$emit('tab-change');
      if (this.activeTab === 1) {
        this.$emit('fetch-logs-data', this.offset);
      }
    }
  },
  methods: {
    prevPage() {
      if (this.offset > 10) {
        this.offset -= 10;
        this.callParentM();
      }
    },
    nextPage() {
      this.offset += 10;
      this.callParentM();
    },
    callParentM() {
      this.$emit('tab-change');
      this.$emit('fetch-logs-data', this.offset);
    },
    showIndividualData(row) {
      this.drawerData = Object.entries(row).map(([key, value]) => ({ key, value }));
      this.drawerVisible = true;
    },
  },
};
</script>

<style scoped>
.button-container {
  position: absolute;
  bottom: 20px;
  right: 20px;
}
.button {
  margin: 10px;
}
</style> -->