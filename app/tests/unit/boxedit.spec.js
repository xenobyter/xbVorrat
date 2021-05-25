import { shallowMount } from '@vue/test-utils'
import boxedit from '@/components/boxedit.vue'

describe('boxedit.vue', () => {
  const wrapper = shallowMount(boxedit, {
    props: { box: { id: 1, name: "PropName", notiz: "PropNotiz" } },
    data() {
      return { res: Object };
    },
  })

  it('Formular wird angezeigt', () => {
    expect(wrapper.text()).toBe('BearbeitenID:Name:Notiz:BearbeitenSchliessen')
  })
  it('Schliessen sendet Event', async () => {
    const btn = wrapper.findAll('button')[1]
    await btn.trigger('click')
    expect(wrapper.emitted().closeedit).toBeTruthy()
  })
  it('Ändern sendet Event mit Daten', async () => {
    const btn = wrapper.findAll('button')[0]
    await btn.trigger('click')
    expect(wrapper.emitted().doedit[0][0]).toEqual({ id: 1, name: "PropName", notiz: "PropNotiz" })
  })
  it('Ändern sendet Event mit neuen Daten', async () => {
    const btn = wrapper.findAll('button')[0]
    const inp = wrapper.findAll('input')[1]
    inp.setValue('ResName')
    await btn.trigger('click')
    expect(wrapper.emitted().doedit[0][0]).toEqual({ id: 1, name: "ResName", notiz: "PropNotiz" })
  })
})
