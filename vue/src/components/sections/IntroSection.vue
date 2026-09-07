<script setup>
import { importIconsProduct } from '@/utils/images'
const iconsProduct = importIconsProduct()

import imgMainPhone from '@/assets/images/banners/main-phone.png'
</script>
<template>
    	<section class="mb-5 mt-5">
	<div class="container">
		<main class="card p-3">
			<div class="row gap-x-3">
				<aside class="lg:w-65">
					<nav class="nav nav-col max-lg:hidden">
						<a class="nav-link" href="#">
							<img :src="iconsProduct.laptop" width="20" height="20" alt="">
							Electronics
						</a>
						<a class="nav-link" href="#">
							<img :src="iconsProduct.tshirt" width="20" height="20" alt="">
							Clothes and wear
						</a>
						<a class="nav-link" href="#">
							<img :src="iconsProduct.sofa" width="20" height="20" alt="">
							Home interiors
						</a>
						<a class="nav-link" href="#">
							<img :src="iconsProduct.book" width="20" height="20" alt="">
							Books & magazines
						</a>
						<a class="nav-link" href="#">
							<img :src="iconsProduct.painting" width="20" height="20" alt="">
							Tools, equipments
						</a>
						<a class="nav-link" href="#">
							<img :src="iconsProduct.ball" width="20" height="20" alt="">
							Sports and outdoor
						</a>
						<a class="nav-link" href="#">
							<img :src="iconsProduct.fish" width="20" height="20" alt="">
							Animal and pets
						</a>
						<a class="nav-link" href="#">
							<img :src="iconsProduct.toy" width="20" height="20" alt="">
							Toys for Kids
						</a>
						<a class="nav-link" href="#">
							<img :src="iconsProduct.more" width="20" height="20" alt="">
							More category
						</a>
					</nav>
				</aside>
				<div class="lg:col">
					<!-- ===== banners =====  -->
					<div class="row gap-3">
						<div class="md:col-9">
							<!-- banner main -->
							<article class="bg-primary-200 relative rounded overflow-hidden">
								<div class="absolute top-10 left-10">
									<h2 class="mb-5 max-sm:text-lg"> <span class="text-primary">New trending</span> <br> Electronic items</h2>
									<a href="#" class="btn btn-neutral btn-lg">Learn more</a>
								</div>
								<img :src="imgMainPhone" alt="Ecommerce banner" class="lg:h-90 w-full  rounded">
							</article>
							<!-- banner main end.// -->
						</div>
						<div class="md:col-3">
							<!-- banner small -->
							<article class="rounded overflow-hidden bg-dark h-full bg-position-center bg-size-cover" style="background-image: url(images/banners/man.png);">
								<div class="bg-alpha-6 py-10 text-white text-center d-flex flex-col justify-content-between h-full">
									<h5 class="text-xl mb-5">Get US $10 off <br> for first order</h5>
									<div>
										<a href="" class="btn btn-outline border-white rounded-full">Get offer</a>
									</div>
								</div>
							</article>
							<!-- banner small end.// -->
						</div>
					</div> <!-- row.// -->
					<!-- ===== banners =====  -->
				</div>
				<!-- col.// -->
			</div>
		</main>
	</div> <!-- container end.// -->
</section>

</template>