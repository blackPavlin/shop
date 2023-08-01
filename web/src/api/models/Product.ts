/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */

import type { ImageList } from './ImageList';

export type Product = {
    id: number;
    categoryId: number;
    name: string;
    description?: string;
    amount: number;
    price: number;
    images: ImageList;
};

