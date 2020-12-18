package team01.yowayowa.sekibutton

import android.content.Intent
import android.os.Bundle
import android.util.Log
import android.widget.Toast
import androidx.preference.PreferenceCategory
import androidx.preference.PreferenceFragmentCompat
import androidx.preference.PreferenceScreen

class PreferenceFragment :PreferenceFragmentCompat(){
    override fun onCreatePreferences(savedInstanceState: Bundle?, rootKey: String?) {
        setPreferencesFromResource(R.xml.preferences,rootKey)
        //デバッグ時のみデバッグの欄を表示させる
        if (BuildConfig.DEBUG) {
            //デバッグ欄の表示
            val signaturePreference: PreferenceCategory? = findPreference("debug")
            signaturePreference?.isVisible = true

            //ウォークスルーの表示を行うクリックリスナー
            val openFirstActivity: PreferenceScreen? = findPreference("openWalkThroughActivity")
            openFirstActivity?.setOnPreferenceClickListener {
                val intent = Intent(activity, WalkThroughActivity::class.java)
                Toast.makeText(activity, "ウォークスルーを強制的に開きます", Toast.LENGTH_SHORT).show()
                startActivity(intent)
                true
            }
        }
    }
}