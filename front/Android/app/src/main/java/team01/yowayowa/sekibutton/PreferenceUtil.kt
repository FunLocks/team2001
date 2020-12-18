package team01.yowayowa.sekibutton

import android.app.Activity
import android.content.Context
import android.content.SharedPreferences.Editor
import androidx.preference.PreferenceManager


class PreferencesUtil {
    //初回起動かどうかの判定
    public fun isFirstJudgment(activity: Activity?): Boolean {
        val sharedPreferences = PreferenceManager.getDefaultSharedPreferences(activity)
        return sharedPreferences?.getBoolean(Key.IS_FIRST.name,true)?:true
    }

    //初回起動フラグの設定
    public fun setFirstFlag(activity: Activity?) {
        PreferenceManager.getDefaultSharedPreferences(activity).edit().apply(){
            putBoolean(Key.IS_FIRST.name,false)
            apply()
        }
    }

    // 設定値の識別子
    private enum class Key {
        IS_FIRST
    }
}