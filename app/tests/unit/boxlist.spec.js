import { shallowMount } from '@vue/test-utils'
import boxlist from '@/components/boxlist.vue'

describe('boxlist.vue', () => {
  const wrapper = shallowMount(boxlist, {
    data() {
      return {
        boxes: [{ id: 1, name: "Name1", notiz: "Notiz1" }],
      }
    }
  })
  it('Überschrift Boxen', () => {
    expect(wrapper.html()).toMatch('<h1>Boxen</h1>')
  })
  it('Testbox 1', () => {
    expect(wrapper.text()).toMatch('Name1Notiz1🖉🗑')
  })
  wrapper.setData({
    boxes: [{ id: 2, name: "Name2", notiz: "Notiz2" }],
  })
  it('Testbox 2', () => {
    expect(wrapper.text()).toMatch('Name2Notiz2🖉🗑')
  })
  it('Starte Bearbeiten für Box1', async (done) => {
    const link = wrapper.findAll('a.edit')[0]
    await link.trigger('click')
    expect(wrapper.vm.showedit).toBe(true)
    done()
  })
})
