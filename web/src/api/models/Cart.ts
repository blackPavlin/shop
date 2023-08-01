/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */

import type { CartProductList } from './CartProductList';

export type Cart = {
    id: string;
    userId: string;
    products: CartProductList;
};

