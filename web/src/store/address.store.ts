import { defineStore } from "pinia";
import { AddressService } from "@/api/services/AddressService";
import type { AddressList, CreateAddressRequest } from "@/api";

type State = {
  addresses: AddressList;
};

export const useAddressStore = defineStore("address", {
  state: (): State => ({
    addresses: [],
  }),
  actions: {
    async loadAddresses(): Promise<void> {
      this.addresses = await AddressService.getAddress();
    },
    async createAddress(body: CreateAddressRequest): Promise<void> {
      await AddressService.postAddress(body);
    },
  },
  getters: {
    getAddresses(): AddressList {
      return this.addresses;
    },
  },
});
