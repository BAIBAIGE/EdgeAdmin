Vue.component("http-cache-refs-box", {
	props: ["v-cache-refs"],
	data: function () {
		let refs = this.vCacheRefs
		if (refs == null) {
			refs = []
		}
		return {
			refs: refs
		}
	},
	methods: {
		timeUnitName: function (unit) {
			switch (unit) {
				case "ms":
					return "毫秒"
				case "second":
					return "秒"
				case "minute":
					return "分钟"
				case "hour":
					return "小时"
				case "day":
					return "天"
				case "week":
					return "周 "
			}
			return unit
		}
	},
	template: `<div>
	<input type="hidden" name="refsJSON" :value="JSON.stringify(refs)"/>
	
	<p class="comment" v-if="refs.length == 0">暂时还没有缓存条件。</p>
	<div v-show="refs.length > 0">
		<table class="ui table selectable celled">
			<thead>
				<tr>
					<th>缓存条件</th>
					<th class="width10">缓存时间</th>
					<th class="two op">操作</th>
				</tr>
				<tr v-for="(cacheRef, index) in refs">
					<td>
						<http-request-conds-view :v-conds="cacheRef.conds"></http-request-conds-view>
					</td>
					<td>{{cacheRef.life.count}} {{timeUnitName(cacheRef.life.unit)}}</td>
					<td>
						<a href="" @click.prevent="updateRef(index, cacheRef)">修改</a> &nbsp;
						<a href="" @click.prevent="removeRef(index)">删除</a>
					</td>
				</tr>
			</thead>
		</table>
	</div>
	<div class="margin"></div>
</div>`
})