import { ref, computed } from 'vue'
import { defineStore } from 'pinia'

export const usePublicStore = defineStore('public', {
    state: () => ({
        title:"",
        breadcrumb:[],
        pageLoading:false,
        requestError:false,
    }),
    getters: {
        GetBreadcrumb(state) {
            return state.breadcrumb
        },
    },
    actions: {
        SetTitle(title) {
            this.title = title;
        },
        SetBreadcrumb(breadcrumb) {
            this.breadcrumb = breadcrumb;
        },
        SetPageLoading(loading) {
            this.pageLoading = loading;
        },
        SetRequestError(error) {
            this.requestError = error;
        }
    },
})
