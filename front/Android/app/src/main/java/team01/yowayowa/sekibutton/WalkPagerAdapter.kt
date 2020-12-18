package team01.yowayowa.sekibutton

import androidx.fragment.app.Fragment
import androidx.fragment.app.FragmentActivity
import androidx.fragment.app.FragmentManager
import androidx.viewpager2.adapter.FragmentStateAdapter

class WalkPagerAdapter(fm: FragmentActivity): FragmentStateAdapter(fm)  {
    private val res : List<Fragment> = listOf<Fragment>(
        WarkFragment1(),
        WarkFragment2(),
        WarkFragment3(),
        WarkFragment4(),
        WarkFragment5(),
        WarkFragment6()
    )

    override fun createFragment(position: Int): Fragment {
        return res[position]
    }

    override fun getItemCount(): Int {
        return res.size
    }

}