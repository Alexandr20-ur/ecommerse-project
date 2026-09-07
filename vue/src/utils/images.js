// utils/images.js
export const importAllImages = () => {
  const images = import.meta.glob(`@/assets/images/*/*.{png,jpg,jpeg,svg,gif,webp}`, { eager: true })
  
  return fromEntries(images)
}

export const importImages = () => {
  const images = import.meta.glob(`@/assets/images/*.{png,jpg,jpeg,svg,gif,webp}`, { eager: true })
  return fromEntries(images)
}


export const importIconsProduct = () => {
    const images = import.meta.glob('@/assets/images/icons-product/*.{png,jpg,jpeg,svg,gif,webp}', { eager: true })
    return fromEntries(images)
}

export const importProductTech = () => {
    const images = import.meta.glob('@/assets/images/product-tech/*.{png,jpg,jpeg,svg,gif,webp}', { eager: true })
    return fromEntries(images)
}

export const importCategory = () => {
    const images = import.meta.glob('@/assets/images/category/*.{png,jpg,jpeg,svg,gif,webp}', { eager: true })
    return fromEntries(images)
}

export const importProductBasic = () => {
    const images = import.meta.glob('@/assets/images/product-basic/*.{png,jpg,jpeg,svg,gif,webp}', { eager: true })
    return fromEntries(images)
}


export const fromEntries = (images) => {
    return Object.fromEntries(
    Object.entries(images).map(([path, module]) => {
      const fileName = path.split('/').pop().replace(/\.[^.]+$/, '')
      return [fileName, path]
    })
  )
}